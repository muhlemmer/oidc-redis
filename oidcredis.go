package oidcredis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/zitadel/oidc/pkg/oidc"
	"github.com/zitadel/oidc/pkg/op"
	"gopkg.in/square/go-jose.v2"
)

// RedisClient describes the methods consumed by this package.
// This allows for switching between singular/cluser/sentinel clients.
type RedisClient interface {
	Close() error
	Ping(context.Context) *redis.StatusCmd
}

var ErrUnimplemented = errors.New("unimplemented method")

type AuthStorage struct {
	AuthRequests RedisClient
	Tokens       RedisClient
}

func (s *AuthStorage) Close() error {
	fcs := [...]func() error{
		s.AuthRequests.Close,
		s.Tokens.Close,
	}

	ec := make(chan error, len(fcs))

	for _, f := range fcs {
		go func(f func() error) {
			ec <- f()
		}(f)
	}

	for i := 0; i < len(fcs); i++ {
		if err := <-ec; err != nil {
			return fmt.Errorf("auth storage close: %w", err)
		}
	}

	return nil
}

type healthCheckFunc func(context.Context) *redis.StatusCmd

func (s *AuthStorage) Health(ctx context.Context) error {
	checks := [...]healthCheckFunc{
		s.AuthRequests.Ping,
		s.Tokens.Ping,
	}

	sc := make(chan *redis.StatusCmd, len(checks))

	for _, check := range checks {
		go func(check healthCheckFunc) {
			sc <- check(ctx)
		}(check)
	}

	for i := 0; i < len(checks); i++ {
		status := <-sc
		if err := status.Err(); err != nil {
			return fmt.Errorf("auth storage health: %w", err)
		}
	}

	return nil
}

func (s *AuthStorage) CreateAuthRequest(context.Context, *oidc.AuthRequest, string) (op.AuthRequest, error) {
	return nil, ErrUnimplemented
}
func (s *AuthStorage) AuthRequestByID(context.Context, string) (op.AuthRequest, error) {
	return nil, ErrUnimplemented
}
func (s *AuthStorage) AuthRequestByCode(context.Context, string) (op.AuthRequest, error) {
	return nil, ErrUnimplemented
}
func (s *AuthStorage) SaveAuthCode(context.Context, string, string) error {
	return ErrUnimplemented
}
func (s *AuthStorage) DeleteAuthRequest(context.Context, string) error {
	return ErrUnimplemented
}

func (s *AuthStorage) CreateAccessToken(context.Context, op.TokenRequest) (accessTokenID string, expiration time.Time, err error) {
	return "", time.Time{}, ErrUnimplemented
}

func (s *AuthStorage) CreateAccessAndRefreshTokens(ctx context.Context, request op.TokenRequest, currentRefreshToken string) (accessTokenID string, newRefreshTokenID string, expiration time.Time, err error) {
	return "", "", time.Time{}, ErrUnimplemented
}

func (s *AuthStorage) TokenRequestByRefreshToken(ctx context.Context, refreshTokenID string) (op.RefreshTokenRequest, error) {
	return nil, ErrUnimplemented
}

func (s *AuthStorage) TerminateSession(ctx context.Context, userID string, clientID string) error {
	return ErrUnimplemented
}

func (s *AuthStorage) RevokeToken(ctx context.Context, tokenID string, userID string, clientID string) *oidc.Error {
	return &oidc.Error{Parent: ErrUnimplemented}
}

func (s *AuthStorage) GetSigningKey(context.Context, chan<- jose.SigningKey) {}

func (s *AuthStorage) GetKeySet(context.Context) (*jose.JSONWebKeySet, error) {
	return nil, ErrUnimplemented
}
