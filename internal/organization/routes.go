package organization

import "net/http"

// RegisterRoutes owns organization endpoints; the composition root supplies the
// existing authentication/session guard. Application authorization stays in the use case.
func RegisterRoutes(mux *http.ServeMux, controller *Controller, protect func(http.Handler) http.Handler) {
	mux.Handle("POST /api/v1/organizations", protect(controller))
}
