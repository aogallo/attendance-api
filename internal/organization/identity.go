package organization

import "context"

type identityKey struct{}

// WithIdentity is called only after trusted authentication and administrative session verification.
func WithIdentity(ctx context.Context, identity Identity) context.Context {
	return context.WithValue(ctx, identityKey{}, identity)
}

// IdentityFromContext rejects absent or incomplete administrative identities.
func IdentityFromContext(ctx context.Context) (Identity, bool) {
	identity, ok := ctx.Value(identityKey{}).(Identity)
	return identity, ok && identity.Issuer != "" && identity.Subject != ""
}
