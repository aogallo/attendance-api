package auth

import (
	"context"
	"errors"
	"reflect"
)

// Authenticator requires both cryptographic identity and live session enforcement.
// It grants no application role. Callers must separately resolve local authority.
type Authenticator struct {
	tokens   TokenVerifier
	sessions SessionVerifier
}

func NewAuthenticator(tokens TokenVerifier, sessions SessionVerifier) (*Authenticator, error) {
	if missing(tokens) || missing(sessions) {
		return nil, errors.New("token and session verifiers are required")
	}
	return &Authenticator{tokens: tokens, sessions: sessions}, nil
}

func missing(v any) bool {
	if v == nil {
		return true
	}
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return value.IsNil()
	}
	return false
}

func (a *Authenticator) Verify(ctx context.Context, raw string) (Principal, error) {
	p, err := a.tokens.Verify(ctx, raw)
	if err == nil {
		err = a.sessions.Verify(ctx, p)
	}
	if err != nil {
		switch {
		case errors.Is(err, ErrUnauthenticated):
			return Principal{}, ErrUnauthenticated
		case errors.Is(err, ErrMFARequired):
			return Principal{}, ErrMFARequired
		default:
			return Principal{}, ErrUnavailable
		}
	}
	return p, nil
}
