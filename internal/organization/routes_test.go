package organization_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aogallo/attendance-api/internal/organization"
)

func TestRegisterRoutesAppliesProtection(t *testing.T) {
	repo := &repository{}
	mux := http.NewServeMux()
	registrations, calls := 0, 0
	protect := func(next http.Handler) http.Handler {
		registrations++
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			calls++
			w.WriteHeader(http.StatusUnauthorized)
		})
	}
	organization.RegisterRoutes(mux, organization.NewController(organization.NewService(repo)), protect)
	req := httptest.NewRequest("POST", "/api/v1/organizations", strings.NewReader(`{"name":"School","initialAdministratorEmail":"a@example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if registrations != 1 || calls != 1 || w.Code != http.StatusUnauthorized || repo.calls != 0 {
		t.Fatalf("registrations=%d guards=%d status=%d repository calls=%d", registrations, calls, w.Code, repo.calls)
	}
}
