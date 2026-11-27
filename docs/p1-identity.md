# P1 administrative identity verification

`internal/auth` verifies signed administrative credentials; it does **not** grant application authority or prove a session remains active. Administrative writes require both token verification and an externally supplied session verifier. P1 must keep writes unavailable until P2 supplies verified session enforcement.

## Integration seam

```go
func NewJWTVerifier(ctx context.Context, config JWTConfig) (*JWTVerifier, error)
func NewAuthenticator(tokens TokenVerifier, sessions SessionVerifier) (*Authenticator, error)

// JWTVerifier and Authenticator both implement TokenVerifier.
type TokenVerifier interface {
    Verify(context.Context, string) (Principal, error)
}
type SessionVerifier interface {
    Verify(context.Context, Principal) error
}
type Principal struct {
    Issuer    string
    Subject   string
    SessionID string
}
type JWTConfig struct {
    Issuer     string
    JWKSURL    string
    Audience   string
    Algorithms []string
}
```

1. Supply operator-controlled issuer and JWKS URLs, explicit audience `authenticated`, and a nonempty asymmetric algorithm allowlist (`ES256`, `RS256`, or `EdDSA`). Both URLs require HTTPS and reject credentials, query strings and fragments. Invalid configuration fails without a request.
2. Pass the process shutdown context to `NewJWTVerifier`; construction is lazy and performs no network request. The constructor has no insecure-development switch or runtime local signing keys. Tests alone use ephemeral keys and a local TLS server's trust store through an unexported constructor.
3. Compose the JWT verifier with a real `SessionVerifier` through `NewAuthenticator`. Nil and typed-nil dependencies are rejected. There is no default passing session verifier. Until P2 exists, executable wiring must fail closed, not call the JWT verifier alone to authorize writes.
4. Extract the raw access token in the transport layer and call the composed verifier. On success, resolve local administrative identity and authority by **issuer plus subject**, not token roles, email, headers, or request-body assertions. This package owns no routes or Bearer parsing.

## Verification and error contract

JWT parsing and signature verification use `github.com/golang-jwt/jwt/v5`; JWKS parsing and key selection use `github.com/MicahParks/keyfunc/v3`. These are community libraries, not official Go or Supabase libraries. The package implements no cryptographic primitive.

| Check | Result |
|---|---|
| Allowed signature algorithm and signature; exact configured issuer and audience; required, unexpired `exp`; UUID-shaped `sub` and `session_id`; nonempty `kid` | Invalid credentials return `ErrUnauthenticated` → HTTP 401 |
| Signed `aal` equals `aal2`, and signed `amr` entries include both `method: password` and `method: totp` | Missing completed assurance returns `ErrMFARequired` → HTTP 403 |
| JWKS HTTP/network/parse failure, canceled context, or unrecognized downstream verifier error | `ErrUnavailable` → HTTP 503 |
| Unknown `kid` in an available JWKS | `ErrUnauthenticated` → HTTP 401 |

Rejected credentials return an empty `Principal`. Exported verification errors are exact, safe sentinels: no raw tokens, claim contents, provider errors, or response bodies. The caller owns generic Problem JSON and must not log credentials. Provider role claims and MFA enrollment metadata confer no authority; only verified signed assurance satisfies this token-level MFA check.

## Network and cache bounds

Only the configured JWKS URL is fetched. Token `jku`/`x5u` values are ignored and redirects are prohibited. Each verification has a three-second network/wait deadline, also constrained by its caller and shutdown contexts. JWKS responses are limited to 1 MiB; raw tokens are limited to 16 KiB. There are no background refresh workers.

A serialized cache keeps one parsed JWKS for five minutes. Unknown key IDs do not trigger refresh, preventing attacker-controlled refresh storms; a new signing key may therefore be rejected for up to five minutes after publication. Failed refreshes back off for five seconds. Once a cache entry expires, failed refreshes return unavailable rather than use stale keys. Already cached keys continue working during an outage until that five-minute expiry, including keys removed upstream during that interval. Operators must allow for this rotation/revocation window. Trust-anchor URLs are privileged configuration, not request input.

## P2 limits and review path

JWT verification alone cannot prove sign-out revocation, factor-recovery invalidation, the approved 30-minute inactivity limit, the eight-hour absolute limit, or fresh password plus TOTP reauthentication after either limit. P2's session verifier must enforce those requirements using trusted session state and explicit activity semantics. Token refresh, AMR evidence and JWT expiry are not substitutes.

No bootstrap, login, session store, provider operations, invitation acceptance, or email-verification proof is implemented. P1 is not invitation acceptance: this principal cannot establish that an invited email is verified. No Supabase project is contacted by these tests or constructors.

Review `jwt_test.go`, `network_test.go`, and `auth_test.go` against [approved implementation plan](../specs/organization-onboarding/implementation-plan.md), [administrative access policy](security/administrative-access.md), and [TD-004](technical-decisions/004-initial-administrative-authentication.md). Tests exercise exported `Verify` with locally signed credentials, safe rejection sentinels, role exclusion, JWKS failures/cache behavior, concurrent verification and mandatory session composition.

```sh
go test ./internal/auth -count=1
go test -race ./internal/auth -count=1
```
