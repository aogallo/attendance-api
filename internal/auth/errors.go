package auth

import "errors"

var (
	ErrUnauthenticated = errors.New("unauthenticated")
	ErrMFARequired     = errors.New("MFA required")
	ErrUnavailable     = errors.New("authentication unavailable")
)
