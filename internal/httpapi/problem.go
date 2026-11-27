// Package httpapi contains shared HTTP transport helpers, not business rules.
package httpapi

import (
	"encoding/json"
	"net/http"
)

func WriteProblem(w http.ResponseWriter, status int, code string, fields any) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.Header().Set("Cache-Control", "no-store")
	if status == http.StatusUnauthorized {
		w.Header().Set("WWW-Authenticate", `Bearer realm="attendance-api"`)
	}
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(Problem{Type: "about:blank", Title: http.StatusText(status), Status: status, Code: code, Errors: fields})
}
