package organization

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime"
	"net/http"

	"github.com/aogallo/attendance-api/internal/httpapi"
)

type Controller struct{ service *Service }

func NewController(
	service *Service,
) *Controller {
	return &Controller{service: service}
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	identity, ok := IdentityFromContext(r.Context())
	if !ok {
		httpapi.WriteProblem(w, 401, "unauthenticated", nil)
		return
	}
	media, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || media != "application/json" {
		httpapi.WriteProblem(w, 415, "unsupported_media_type", nil)
		return
	}
	body, err := io.ReadAll(http.MaxBytesReader(w, r.Body, maxRequestBodyBytes))
	if err != nil {
		var large *http.MaxBytesError
		if errors.As(err, &large) {
			httpapi.WriteProblem(w, 413, "request_too_large", nil)
		} else {
			httpapi.WriteProblem(w, 400, "invalid_json", nil)
		}
		return
	}
	var request *struct {
		Name                      string `json:"name"`
		InitialAdministratorEmail string `json:"initialAdministratorEmail"`
	}
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&request) != nil || request == nil ||
		decoder.Decode(new(any)) != io.EOF {
		httpapi.WriteProblem(w, 400, "invalid_json", nil)
		return
	}
	result, err := c.service.Create(
		r.Context(),
		identity,
		request.Name,
		request.InitialAdministratorEmail,
	)
	if err != nil {
		var validation *ValidationError
		switch {
		case errors.As(err, &validation):
			httpapi.WriteProblem(w, 422, "validation_failed", validation.Fields)
		case errors.Is(err, ErrForbidden):
			httpapi.WriteProblem(w, 403, "forbidden", nil)
		case errors.Is(err, ErrUnavailable):
			httpapi.WriteProblem(w, 503, "dependency_unavailable", nil)
		default:
			httpapi.WriteProblem(w, 500, "internal_error", nil)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(result)
}
