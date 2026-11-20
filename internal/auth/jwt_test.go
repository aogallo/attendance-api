package auth

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const subject = "12345678-1234-1234-1234-123456789abc"
const session = "abcdefab-1234-1234-1234-123456789abc"

func fixture(t *testing.T) (*JWTVerifier, *ecdsa.PrivateKey, *atomic.Int32) {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	var status atomic.Int32
	status.Store(http.StatusOK)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/keys" {
			t.Errorf("unexpected request: %s", r.URL.Path)
		}
		w.WriteHeader(int(status.Load()))
		_ = json.NewEncoder(w).Encode(map[string]any{"keys": []any{map[string]any{
			"kty": "EC", "crv": "P-256", "use": "sig", "alg": "ES256", "kid": "known",
			"x": base64.RawURLEncoding.EncodeToString(key.X.FillBytes(make([]byte, 32))),
			"y": base64.RawURLEncoding.EncodeToString(key.Y.FillBytes(make([]byte, 32))),
		}}})
	}))
	t.Cleanup(server.Close)
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	v, err := newJWTVerifier(ctx, JWTConfig{
		Issuer: "https://trusted.example/auth/v1", JWKSURL: server.URL + "/keys",
		Audience: "authenticated", Algorithms: []string{"ES256"},
	}, server.Client())
	if err != nil {
		t.Fatal(err)
	}
	return v, key, &status
}

func claims() jwt.MapClaims {
	return jwt.MapClaims{
		"iss": "https://trusted.example/auth/v1", "aud": "authenticated",
		"exp": time.Now().Add(time.Hour).Unix(), "sub": subject, "session_id": session,
		"aal": "aal2", "amr": []any{map[string]any{"method": "password"}, map[string]any{"method": "totp"}},
	}
}

func sign(t *testing.T, key *ecdsa.PrivateKey, c jwt.MapClaims, edit func(*jwt.Token)) string {
	t.Helper()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	token.Header["kid"] = "known"
	if edit != nil {
		edit(token)
	}
	var signingKey any = key
	if token.Method == jwt.SigningMethodHS256 {
		signingKey = []byte("untrusted secret")
	}
	if token.Method == jwt.SigningMethodNone {
		signingKey = jwt.UnsafeAllowNoneSignatureType
	}
	raw, err := token.SignedString(signingKey)
	if err != nil {
		t.Fatal(err)
	}
	return raw
}

func TestVerify(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(jwt.MapClaims)
		edit   func(*jwt.Token)
		want   error
	}{
		{name: "valid"},
		{name: "role is not authority", mutate: func(c jwt.MapClaims) {
			c["role"] = "platform_administrator"
			c["app_metadata"] = map[string]any{"role": "platform_administrator"}
		}},
		{name: "wrong issuer", mutate: func(c jwt.MapClaims) { c["iss"] = "https://attacker.example" }, want: ErrUnauthenticated},
		{name: "wrong audience", mutate: func(c jwt.MapClaims) { c["aud"] = "anon" }, want: ErrUnauthenticated},
		{name: "missing audience", mutate: func(c jwt.MapClaims) { delete(c, "aud") }, want: ErrUnauthenticated},
		{name: "expired", mutate: func(c jwt.MapClaims) { c["exp"] = time.Now().Add(-time.Minute).Unix() }, want: ErrUnauthenticated},
		{name: "missing expiry", mutate: func(c jwt.MapClaims) { delete(c, "exp") }, want: ErrUnauthenticated},
		{name: "missing subject", mutate: func(c jwt.MapClaims) { delete(c, "sub") }, want: ErrUnauthenticated},
		{name: "invalid subject", mutate: func(c jwt.MapClaims) { c["sub"] = "not-a-uuid" }, want: ErrUnauthenticated},
		{name: "missing session", mutate: func(c jwt.MapClaims) { delete(c, "session_id") }, want: ErrUnauthenticated},
		{name: "invalid session", mutate: func(c jwt.MapClaims) { c["session_id"] = "not-a-uuid" }, want: ErrUnauthenticated},
		{name: "password only", mutate: func(c jwt.MapClaims) {
			c["aal"] = "aal1"
			c["amr"] = []any{map[string]any{"method": "password"}}
		}, want: ErrMFARequired},
		{name: "enrolled not verified", mutate: func(c jwt.MapClaims) {
			c["amr"] = []any{map[string]any{"method": "password"}}
			c["factors"] = []string{"totp"}
		}, want: ErrMFARequired},
		{name: "no password", mutate: func(c jwt.MapClaims) { c["amr"] = []any{map[string]any{"method": "totp"}} }, want: ErrMFARequired},
		{name: "missing assurance", mutate: func(c jwt.MapClaims) { delete(c, "aal") }, want: ErrMFARequired},
		{name: "unknown kid", edit: func(t *jwt.Token) { t.Header["kid"] = "unknown" }, want: ErrUnauthenticated},
		{name: "missing kid", edit: func(t *jwt.Token) { delete(t.Header, "kid") }, want: ErrUnauthenticated},
		{name: "token directed URLs ignored", edit: func(t *jwt.Token) {
			t.Header["jku"] = "https://untrusted.invalid/keys"
			t.Header["x5u"] = "https://untrusted.invalid/certificate"
		}},
		{name: "missing issuer", mutate: func(c jwt.MapClaims) { delete(c, "iss") }, want: ErrUnauthenticated},
		{name: "missing AMR", mutate: func(c jwt.MapClaims) { delete(c, "amr") }, want: ErrMFARequired},
		{name: "symmetric algorithm", edit: func(t *jwt.Token) { t.Method = jwt.SigningMethodHS256 }, want: ErrUnauthenticated},
		{name: "unsigned", edit: func(t *jwt.Token) { t.Method = jwt.SigningMethodNone }, want: ErrUnauthenticated},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, key, _ := fixture(t)
			c := claims()
			if tt.mutate != nil {
				tt.mutate(c)
			}
			raw := sign(t, key, c, tt.edit)
			got, err := v.Verify(context.Background(), raw)
			assertResult(t, got, err, tt.want)
			if err != nil && strings.Contains(err.Error(), raw) {
				t.Fatal("error leaked token")
			}
		})
	}
}

func assertResult(t *testing.T, got Principal, err, wantErr error) {
	t.Helper()
	if err != wantErr {
		t.Fatalf("error = %v, want exact sentinel %v", err, wantErr)
	}
	want := Principal{}
	if wantErr == nil {
		want = Principal{Issuer: "https://trusted.example/auth/v1", Subject: subject, SessionID: session}
	}
	if got != want {
		t.Fatalf("principal = %+v, want %+v", got, want)
	}
}

func TestInvalidCredentialsAndOutages(t *testing.T) {
	for _, name := range []string{"bad signature", "missing algorithm", "malformed", "outage", "cached outage", "canceled"} {
		t.Run(name, func(t *testing.T) {
			v, key, status := fixture(t)
			raw := sign(t, key, claims(), nil)
			want := ErrUnauthenticated
			ctx := context.Background()
			switch name {
			case "bad signature":
				other, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
				if err != nil {
					t.Fatal(err)
				}
				raw = sign(t, other, claims(), nil)
			case "missing algorithm":
				parts := strings.Split(raw, ".")
				parts[0] = base64.RawURLEncoding.EncodeToString([]byte(`{"kid":"known"}`))
				raw = strings.Join(parts, ".")
			case "malformed":
				raw = "not a jwt"
			case "outage":
				status.Store(503)
				want = ErrUnavailable
			case "cached outage":
				if _, err := v.Verify(ctx, raw); err != nil {
					t.Fatal(err)
				}
				status.Store(503)
				want = nil
			case "canceled":
				var cancel context.CancelFunc
				ctx, cancel = context.WithCancel(ctx)
				cancel()
				want = ErrUnavailable
			}
			got, err := v.Verify(ctx, raw)
			assertResult(t, got, err, want)
		})
	}
}

func TestUnsafeConfig(t *testing.T) {
	for _, name := range []string{"http", "empty issuer", "empty audience", "wrong audience", "no algorithms", "symmetric", "URL credentials"} {
		t.Run(name, func(t *testing.T) {
			c := JWTConfig{Issuer: "https://trusted.example", JWKSURL: "https://trusted.example/keys", Audience: "authenticated", Algorithms: []string{"ES256"}}
			switch name {
			case "http":
				c.JWKSURL = "http://localhost/keys"
			case "empty issuer":
				c.Issuer = ""
			case "empty audience":
				c.Audience = ""
			case "wrong audience":
				c.Audience = "anon"
			case "no algorithms":
				c.Algorithms = nil
			case "symmetric":
				c.Algorithms = []string{"HS256"}
			case "URL credentials":
				c.JWKSURL = "https://user:password@trusted.example/keys"
			}
			if _, err := NewJWTVerifier(context.Background(), c); err == nil {
				t.Fatal("unsafe configuration accepted")
			}
		})
	}
}

func TestConcurrentVerification(t *testing.T) {
	v, key, _ := fixture(t)
	raw := sign(t, key, claims(), nil)
	var workers sync.WaitGroup
	for range 16 {
		workers.Go(func() {
			got, err := v.Verify(context.Background(), raw)
			if err != nil || got != (Principal{Issuer: "https://trusted.example/auth/v1", Subject: subject, SessionID: session}) {
				t.Errorf("concurrent verification: principal=%+v error=%v", got, err)
			}
		})
	}
	workers.Wait()
}
