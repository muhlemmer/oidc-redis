package oidcredis

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/zitadel/oidc/pkg/oidc"
	"github.com/zitadel/oidc/pkg/op"
	"gopkg.in/square/go-jose.v2"
)

type AuthStorage interface {
	op.AuthStorage
	Health(ctx context.Context) error
}

var ErrUnimplemented = errors.New("unimplemented method")

// increments to the Redis DB number
const (
	DBAuthRequests = iota
	DBTokens
)

type authStorage struct {
	auth   *redis.Client
	tokens *redis.Client
}

// NewAuthStorage connects to a single redis server.
// It currently consumes 2 databases:
//
//	opts.DB + 0: AuthRequests
//	opts.DB + 1: Tokens
func NewAuthStorage(opts redis.Options) AuthStorage {
	authOpts := opts
	authOpts.TLSConfig = opts.TLSConfig.Clone()
	authOpts.DB += DBAuthRequests

	tokenOpts := opts
	tokenOpts.TLSConfig = opts.TLSConfig.Clone()
	tokenOpts.DB += DBTokens

	return &authStorage{
		auth:   redis.NewClient(&authOpts),
		tokens: redis.NewClient(&authOpts),
	}
}

func (s *authStorage) Health(ctx context.Context) error {
	return ErrUnimplemented
}

func (s *authStorage) CreateAuthRequest(context.Context, *oidc.AuthRequest, string) (op.AuthRequest, error) {
	return nil, ErrUnimplemented
}
func (s *authStorage) AuthRequestByID(context.Context, string) (op.AuthRequest, error) {
	return nil, ErrUnimplemented
}
func (s *authStorage) AuthRequestByCode(context.Context, string) (op.AuthRequest, error) {
	return nil, ErrUnimplemented
}
func (s *authStorage) SaveAuthCode(context.Context, string, string) error {
	return ErrUnimplemented
}
func (s *authStorage) DeleteAuthRequest(context.Context, string) error {
	return ErrUnimplemented
}

func (s *authStorage) CreateAccessToken(context.Context, op.TokenRequest) (accessTokenID string, expiration time.Time, err error) {
	return "", time.Time{}, ErrUnimplemented
}

func (s *authStorage) CreateAccessAndRefreshTokens(ctx context.Context, request op.TokenRequest, currentRefreshToken string) (accessTokenID string, newRefreshTokenID string, expiration time.Time, err error) {
	return "", "", time.Time{}, ErrUnimplemented
}

func (s *authStorage) TokenRequestByRefreshToken(ctx context.Context, refreshTokenID string) (op.RefreshTokenRequest, error) {
	return nil, ErrUnimplemented
}

func (s *authStorage) TerminateSession(ctx context.Context, userID string, clientID string) error {
	return ErrUnimplemented
}

func (s *authStorage) RevokeToken(ctx context.Context, tokenID string, userID string, clientID string) *oidc.Error {
	return &oidc.Error{Parent: ErrUnimplemented}
}

func (s *authStorage) GetSigningKey(context.Context, chan<- jose.SigningKey) {}

func (s *authStorage) GetKeySet(context.Context) (*jose.JSONWebKeySet, error) {
	return nil, ErrUnimplemented
}
