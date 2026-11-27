package auth

import "context"

// Principal deliberately contains neither email nor application roles.
type Principal struct {
	Issuer    string
	Subject   string
	SessionID string
}

type TokenVerifier interface {
	Verify(context.Context, string) (Principal, error)
}

// SessionVerifier must enforce revocation, inactivity and absolute session limits.
type SessionVerifier interface {
	Verify(context.Context, Principal) error
}

// JWTConfig contains operator-controlled trust anchors, never request values.
// Audience must be "authenticated". Algorithms may contain ES256, RS256 or EdDSA.
type JWTConfig struct {
	Issuer     string
	JWKSURL    string
	Audience   string
	Algorithms []string
}
