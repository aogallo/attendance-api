package auth

import "time"

const (
	verificationTimeout = 3 * time.Second
	maxTokenBytes       = 16 * 1024
	maxJWKSBytes        = 1024 * 1024
	jwksCacheLifetime   = 5 * time.Minute
	jwksRetryBackoff    = 5 * time.Second
)
