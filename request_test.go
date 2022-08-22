package oidcredis

import "github.com/zitadel/oidc/pkg/op"

// compile-time implementation checks
var (
	_ op.AuthRequest         = new(Request)
	_ op.RefreshTokenRequest = new(Request)
)
