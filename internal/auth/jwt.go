// Package auth verifies administrative identity, not application authority.
package auth

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"slices"
	"time"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
)

// JWTVerifier uses a bounded, lazy JWKS cache. Cancel its constructor context
// on shutdown. There are no background refresh goroutines or startup requests.
type JWTVerifier struct {
	lifetime   context.Context
	config     JWTConfig
	client     *http.Client
	gate       chan struct{}
	keys       keyfunc.Keyfunc
	expires    time.Time
	retryAfter time.Time
}

// NewJWTVerifier accepts HTTPS trust anchors only and makes no network calls.
func NewJWTVerifier(ctx context.Context, c JWTConfig) (*JWTVerifier, error) {
	return newJWTVerifier(ctx, c, &http.Client{})
}

// Tests can supply a TLS httptest client's trust store, not disable URL checks.
func newJWTVerifier(ctx context.Context, c JWTConfig, client *http.Client) (*JWTVerifier, error) {
	if ctx == nil || !safeURL(c.Issuer) || !safeURL(c.JWKSURL) || c.Audience != "authenticated" || len(c.Algorithms) == 0 {
		return nil, errors.New("invalid JWT trust configuration")
	}
	for _, alg := range c.Algorithms {
		if !slices.Contains([]string{"ES256", "RS256", "EdDSA"}, alg) {
			return nil, errors.New("invalid JWT algorithm configuration")
		}
	}
	c.Algorithms = slices.Clone(c.Algorithms)
	bounded := *client
	bounded.Timeout = verificationTimeout
	bounded.CheckRedirect = func(*http.Request, []*http.Request) error {
		return errors.New("JWKS redirects prohibited")
	}
	return &JWTVerifier{lifetime: ctx, config: c, client: &bounded, gate: make(chan struct{}, 1)}, nil
}

func safeURL(raw string) bool {
	u, err := url.Parse(raw)
	return err == nil && u.Scheme == "https" && u.Hostname() != "" && u.User == nil && u.RawQuery == "" && u.Fragment == ""
}

var uuidShape = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

type accessClaims struct {
	jwt.RegisteredClaims
	SessionID string `json:"session_id"`
	AAL       string `json:"aal"`
	AMR       []struct {
		Method string `json:"method"`
	} `json:"amr"`
}

func (v *JWTVerifier) Verify(ctx context.Context, raw string) (Principal, error) {
	if ctx.Err() != nil || v.lifetime.Err() != nil {
		return Principal{}, ErrUnavailable
	}
	if len(raw) == 0 || len(raw) > maxTokenBytes {
		return Principal{}, ErrUnauthenticated
	}
	ctx, cancel := context.WithTimeout(ctx, verificationTimeout)
	defer cancel()
	stop := context.AfterFunc(v.lifetime, cancel)
	defer stop()
	claims := &accessClaims{}
	token, err := jwt.ParseWithClaims(raw, claims, func(token *jwt.Token) (any, error) {
		kid, ok := token.Header["kid"].(string)
		if !ok || kid == "" {
			return nil, ErrUnauthenticated
		}
		keys, err := v.keyset(ctx)
		if err != nil {
			return nil, err
		}
		return keys.KeyfuncCtx(ctx)(token)
	}, jwt.WithValidMethods(v.config.Algorithms), jwt.WithIssuer(v.config.Issuer),
		jwt.WithAudience(v.config.Audience), jwt.WithExpirationRequired())
	if err != nil {
		if errors.Is(err, ErrUnavailable) {
			return Principal{}, ErrUnavailable
		}
		return Principal{}, ErrUnauthenticated
	}
	if ctx.Err() != nil {
		return Principal{}, ErrUnavailable
	}
	if !token.Valid || !uuidShape.MatchString(claims.Subject) || !uuidShape.MatchString(claims.SessionID) {
		return Principal{}, ErrUnauthenticated
	}
	password, totp := false, false
	for _, method := range claims.AMR {
		password = password || method.Method == "password"
		totp = totp || method.Method == "totp"
	}
	if claims.AAL != "aal2" || !password || !totp {
		return Principal{}, ErrMFARequired
	}
	return Principal{Issuer: claims.Issuer, Subject: claims.Subject, SessionID: claims.SessionID}, nil
}

func (v *JWTVerifier) keyset(ctx context.Context) (keyfunc.Keyfunc, error) {
	select {
	case v.gate <- struct{}{}:
	case <-ctx.Done():
		return nil, ErrUnavailable
	}
	defer func() { <-v.gate }()
	now := time.Now()
	if now.Before(v.expires) {
		return v.keys, nil
	}
	if now.Before(v.retryAfter) {
		return nil, ErrUnavailable
	}
	keys, err := v.fetch(ctx)
	if err != nil {
		v.retryAfter = time.Now().Add(jwksRetryBackoff)
		return nil, ErrUnavailable
	}
	v.keys = keys
	v.expires = time.Now().Add(jwksCacheLifetime)
	return keys, nil
}

func (v *JWTVerifier) fetch(ctx context.Context) (keyfunc.Keyfunc, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, v.config.JWKSURL, nil)
	if err != nil {
		return nil, err
	}
	response, err := v.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = response.Body.Close() }()
	if response.StatusCode != http.StatusOK {
		return nil, ErrUnavailable
	}
	body, err := io.ReadAll(io.LimitReader(response.Body, maxJWKSBytes+1))
	if err != nil || len(body) > maxJWKSBytes {
		return nil, ErrUnavailable
	}
	return keyfunc.NewJWKSetJSON(body)
}
