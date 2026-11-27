package organization_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/aogallo/attendance-api/internal/organization"
)

type repository struct {
	calls       int
	name, email string
	err         error
	identity    organization.Identity
}

func (r *repository) Create(_ context.Context, identity organization.Identity, name, email string) (organization.Organization, error) {
	r.calls++
	r.identity = identity
	r.name, r.email = name, email
	if r.err != nil {
		return organization.Organization{}, r.err
	}
	return organization.Organization{ID: "c4f37cde-8880-4bcd-a904-98fda117e8c3", Name: name, InitialAdministratorEmail: email, CreatedAt: time.Date(2026, 11, 5, 14, 30, 0, 0, time.UTC)}, nil
}
func TestCreateOrganization(t *testing.T) {
	tests := []struct {
		name, body, contentType string
		status                  int
		calls                   int
		code                    string
	}{
		{"success", `{"name":"  Colegio  ","initialAdministratorEmail":" User+tag@EXAMPLE.COM "}`, "application/json", 201, 1, ""},
		{"missing name", `{"initialAdministratorEmail":"a@example.com"}`, "application/json", 422, 0, "validation_failed"},
		{"missing email", `{"name":"School"}`, "application/json", 422, 0, "validation_failed"},
		{"blank name", `{"name":" \t ","initialAdministratorEmail":"a@example.com"}`, "application/json", 422, 0, "validation_failed"},
		{"long name", `{"name":"` + strings.Repeat("ñ", 201) + `","initialAdministratorEmail":"a@example.com"}`, "application/json", 422, 0, "validation_failed"},
		{"unicode boundary", `{"name":"` + strings.Repeat("ñ", 200) + `","initialAdministratorEmail":"a@example.com"}`, "application/json", 201, 1, ""},
		{"email byte boundary", `{"name":"School","initialAdministratorEmail":"` + strings.Repeat("a", 64) + "@" + strings.Repeat("b", 185) + `.com"}`, "application/json", 201, 1, ""},
		{"email too long", `{"name":"School","initialAdministratorEmail":"` + strings.Repeat("a", 65) + "@" + strings.Repeat("b", 185) + `.com"}`, "application/json", 422, 0, "validation_failed"},
		{"body boundary", `{"name":"School","initialAdministratorEmail":"a@example.com"}` + strings.Repeat(" ", 16384-len(`{"name":"School","initialAdministratorEmail":"a@example.com"}`)), "application/json", 201, 1, ""},
		{"invalid email", `{"name":"School","initialAdministratorEmail":"not-email"}`, "application/json", 422, 0, "validation_failed"},
		{"display name", `{"name":"School","initialAdministratorEmail":"Jane <a@example.com>"}`, "application/json", 422, 0, "validation_failed"},
		{"email list", `{"name":"School","initialAdministratorEmail":"a@example.com,b@example.com"}`, "application/json", 422, 0, "validation_failed"},
		{"unknown field", `{"name":"School","initialAdministratorEmail":"a@example.com","role":"platform_administrator"}`, "application/json", 400, 0, "invalid_json"},
		{"malformed", `{`, "application/json", 400, 0, "invalid_json"},
		{"trailing JSON", `{} {}`, "application/json", 400, 0, "invalid_json"},
		{"null", `null`, "application/json", 400, 0, "invalid_json"},
		{"wrong type", `{"name":42}`, "application/json", 400, 0, "invalid_json"},
		{"unsupported media", `{}`, "text/plain", 415, 0, "unsupported_media_type"},
		{"missing media", `{}`, "", 415, 0, "unsupported_media_type"},
		{"too large", strings.Repeat(" ", 16385), "application/json", 413, 0, "request_too_large"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &repository{}
			handler := organization.NewController(organization.NewService(repo))
			req := httptest.NewRequest(http.MethodPost, "/api/v1/organizations", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", tt.contentType)
			req = req.WithContext(organization.WithIdentity(req.Context(), organization.Identity{Issuer: "https://issuer.example", Subject: "subject"}))
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)
			if w.Code != tt.status || repo.calls != tt.calls {
				t.Fatalf("status=%d calls=%d body=%s", w.Code, repo.calls, w.Body.String())
			}
			if tt.code != "" && !strings.Contains(w.Body.String(), `"code":"`+tt.code+`"`) {
				t.Fatalf("body=%s", w.Body.String())
			}
			if w.Header().Get("Cache-Control") != "no-store" {
				t.Fatal("missing no-store")
			}
			if w.Header().Get("Location") != "" {
				t.Fatal("unimplemented Location")
			}
			if tt.name == "success" {
				if repo.identity != (organization.Identity{Issuer: "https://issuer.example", Subject: "subject"}) {
					t.Fatalf("repository identity=%+v", repo.identity)
				}
				const expected = "{\"id\":\"c4f37cde-8880-4bcd-a904-98fda117e8c3\",\"name\":\"Colegio\",\"initialAdministratorEmail\":\"User+tag@example.com\",\"attendanceConfiguration\":null,\"createdAt\":\"2026-11-05T14:30:00Z\"}\n"
				if w.Body.String() != expected {
					t.Fatalf("body=%q", w.Body.String())
				}
			}
		})
	}
}
func TestRejectNULName(t *testing.T) {
	for _, name := range []string{`A\u0000B`, `\u0000AB`, `AB\u0000`, `\u0000`} {
		t.Run(name, func(t *testing.T) {
			repo := &repository{name: "unchanged", email: "unchanged"}
			h := organization.NewController(organization.NewService(repo))
			req := httptest.NewRequest(http.MethodPost, "/api/v1/organizations", strings.NewReader(`{"name":"`+name+`","initialAdministratorEmail":"a@example.com"}`))
			req.Header.Set("Content-Type", "application/json")
			req = req.WithContext(organization.WithIdentity(req.Context(), organization.Identity{Issuer: "issuer", Subject: "subject"}))
			w := httptest.NewRecorder()
			h.ServeHTTP(w, req)
			const expected = "{\"type\":\"about:blank\",\"title\":\"Unprocessable Entity\",\"status\":422,\"code\":\"validation_failed\",\"errors\":[{\"field\":\"name\",\"code\":\"invalid\"}]}\n"
			if w.Code != http.StatusUnprocessableEntity || w.Body.String() != expected {
				t.Fatalf("status=%d body=%q", w.Code, w.Body.String())
			}
			if w.Header().Get("Content-Type") != "application/problem+json" || w.Header().Get("Cache-Control") != "no-store" || w.Header().Get("Location") != "" {
				t.Fatalf("headers=%v", w.Header())
			}
			if repo.calls != 0 || repo.name != "unchanged" || repo.email != "unchanged" || repo.identity != (organization.Identity{}) {
				t.Fatalf("rejected request changed repository: %+v", repo)
			}
		})
	}
}

func TestRepositoryFailures(t *testing.T) {
	for _, tt := range []struct {
		name   string
		err    error
		status int
		code   string
	}{
		{"no role", organization.ErrForbidden, 403, "forbidden"},
		{"database unavailable", organization.ErrUnavailable, 503, "dependency_unavailable"},
		{"unexpected", errors.New("secret database connection"), 500, "internal_error"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			repo := &repository{err: tt.err}
			h := organization.NewController(organization.NewService(repo))
			req := httptest.NewRequest("POST", "/api/v1/organizations", strings.NewReader(`{"name":"School","initialAdministratorEmail":"a@example.com"}`))
			req.Header.Set("Content-Type", "application/json")
			req = req.WithContext(organization.WithIdentity(req.Context(), organization.Identity{Issuer: "issuer", Subject: "subject"}))
			w := httptest.NewRecorder()
			h.ServeHTTP(w, req)
			if w.Code != tt.status || !strings.Contains(w.Body.String(), `"code":"`+tt.code+`"`) || strings.Contains(w.Body.String(), "secret") {
				t.Fatalf("status=%d body=%s", w.Code, w.Body.String())
			}
		})
	}
}
func TestIdentityContextContract(t *testing.T) {
	valid := organization.Identity{Issuer: "https://issuer.example", Subject: "subject"}
	for _, tt := range []struct {
		name     string
		identity organization.Identity
		present  bool
	}{
		{"empty", organization.Identity{}, false},
		{"missing issuer", organization.Identity{Subject: "subject"}, false},
		{"missing subject", organization.Identity{Issuer: "issuer"}, false},
		{"valid", valid, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := organization.WithIdentity(context.Background(), tt.identity)
			got, ok := organization.IdentityFromContext(ctx)
			if got != tt.identity || ok != tt.present {
				t.Fatalf("identity=%+v present=%v", got, ok)
			}
		})
	}
	if got, ok := organization.IdentityFromContext(context.Background()); ok || got != (organization.Identity{}) {
		t.Fatalf("absent identity=%+v present=%v", got, ok)
	}
}

func TestMissingIdentity(t *testing.T) {
	for _, tt := range []struct {
		name string
		ctx  context.Context
	}{
		{"absent", context.Background()},
		{"empty", organization.WithIdentity(context.Background(), organization.Identity{})},
		{"missing issuer", organization.WithIdentity(context.Background(), organization.Identity{Subject: "subject"})},
		{"missing subject", organization.WithIdentity(context.Background(), organization.Identity{Issuer: "issuer"})},
	} {
		t.Run(tt.name, func(t *testing.T) {
			repo := &repository{}
			service := organization.NewService(repo)
			h := organization.NewController(service)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/organizations", strings.NewReader("{}"))
			h.ServeHTTP(w, req.WithContext(tt.ctx))
			const expected = "{\"type\":\"about:blank\",\"title\":\"Unauthorized\",\"status\":401,\"code\":\"unauthenticated\"}\n"
			if w.Code != 401 || w.Body.String() != expected || repo.calls != 0 {
				t.Fatalf("status=%d calls=%d body=%q", w.Code, repo.calls, w.Body.String())
			}
			if w.Header().Get("Cache-Control") != "no-store" || w.Header().Get("WWW-Authenticate") != `Bearer realm="attendance-api"` {
				t.Fatal("missing authentication response headers")
			}
			identity, _ := organization.IdentityFromContext(tt.ctx)
			result, err := service.Create(tt.ctx, identity, "School", "a@example.com")
			if !errors.Is(err, organization.ErrForbidden) || result.ID != "" || repo.calls != 0 {
				t.Fatalf("result=%+v error=%v calls=%d", result, err, repo.calls)
			}
		})
	}
}
