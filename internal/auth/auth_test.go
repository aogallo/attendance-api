package auth

import (
	"context"
	"errors"
	"testing"
)

type tokenStub struct{ err error }

func (s tokenStub) Verify(context.Context, string) (Principal, error) {
	if s.err != nil {
		return Principal{}, s.err
	}
	return Principal{Issuer: "issuer", Subject: subject, SessionID: session}, nil
}

type sessionStub struct {
	err   error
	calls int
	got   Principal
}

func (s *sessionStub) Verify(_ context.Context, p Principal) error {
	s.calls++
	s.got = p
	return s.err
}

func TestAuthenticator(t *testing.T) {
	for _, tt := range []struct {
		name                       string
		tokenErr, sessionErr, want error
		calls                      int
	}{
		{name: "verified session", calls: 1},
		{name: "revoked", sessionErr: ErrUnauthenticated, want: ErrUnauthenticated, calls: 1},
		{name: "MFA", tokenErr: ErrMFARequired, want: ErrMFARequired},
		{name: "session outage", sessionErr: ErrUnavailable, want: ErrUnavailable, calls: 1},
		{name: "unexpected session error sanitized", sessionErr: errors.New("private provider detail"), want: ErrUnavailable, calls: 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			sessions := &sessionStub{err: tt.sessionErr}
			a, err := NewAuthenticator(tokenStub{tt.tokenErr}, sessions)
			if err != nil {
				t.Fatal(err)
			}
			got, err := a.Verify(context.Background(), "credential")
			if err != tt.want {
				t.Fatalf("error = %v, want exact sentinel %v", err, tt.want)
			}
			want := Principal{}
			if tt.want == nil {
				want = Principal{Issuer: "issuer", Subject: subject, SessionID: session}
			}
			if got != want || sessions.calls != tt.calls {
				t.Fatalf("principal=%+v calls=%d", got, sessions.calls)
			}
			if tt.calls > 0 && sessions.got != (Principal{Issuer: "issuer", Subject: subject, SessionID: session}) {
				t.Fatalf("session received %+v", sessions.got)
			}
		})
	}
	var typedNil *sessionStub
	for _, tt := range []struct {
		tokens   TokenVerifier
		sessions SessionVerifier
	}{{nil, &sessionStub{}}, {tokenStub{}, nil}, {tokenStub{}, typedNil}} {
		if _, err := NewAuthenticator(tt.tokens, tt.sessions); err == nil {
			t.Fatal("missing verifier accepted")
		}
	}
}
