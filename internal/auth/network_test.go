package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
)

func TestBoundedJWKSFailures(t *testing.T) {
	for _, name := range []string{"redirect", "malformed", "oversized", "connection lost"} {
		t.Run(name, func(t *testing.T) {
			_, key, _ := fixture(t)
			var requests atomic.Int32
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requests.Add(1)
				switch name {
				case "redirect":
					w.Header().Set("Location", "https://untrusted.invalid/keys")
					w.WriteHeader(http.StatusFound)
				case "malformed":
					_, _ = w.Write([]byte("not JSON"))
				case "oversized":
					_, _ = w.Write([]byte(strings.Repeat(" ", 1024*1024+1)))
				}
			}))
			t.Cleanup(server.Close)
			ctx, cancel := context.WithCancel(context.Background())
			t.Cleanup(cancel)
			v, err := newJWTVerifier(ctx, JWTConfig{
				Issuer: "https://trusted.example/auth/v1", JWKSURL: server.URL,
				Audience: "authenticated", Algorithms: []string{"ES256"},
			}, server.Client())
			if err != nil {
				t.Fatal(err)
			}
			if requests.Load() != 0 {
				t.Fatal("constructor made a network request")
			}
			if name == "connection lost" {
				server.Close()
			}
			raw := sign(t, key, claims(), nil)
			for range 2 {
				got, err := v.Verify(context.Background(), raw)
				assertResult(t, got, err, ErrUnavailable)
			}
			wantRequests := int32(1)
			if name == "connection lost" {
				wantRequests = 0
			}
			if requests.Load() != wantRequests {
				t.Fatalf("requests = %d, want %d (failure backoff)", requests.Load(), wantRequests)
			}
			cancel()
			got, err := v.Verify(context.Background(), raw)
			assertResult(t, got, err, ErrUnavailable)
		})
	}
}
