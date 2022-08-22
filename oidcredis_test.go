package oidcredis

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/zitadel/oidc/pkg/op"
)

// compile-time implementation checks
var (
	_ op.AuthStorage = new(AuthStorage)
)

var (
	testAuthStorage AuthStorage
	testCTX         context.Context
	errCTX          context.Context
)

func connectRedis() AuthStorage {
	return AuthStorage{
		AuthRequests: redis.NewClient(&redis.Options{
			Addr: "redis:6379",
			DB:   0,
		}),
		Tokens: redis.NewClient(&redis.Options{
			Addr: "redis:6379",
			DB:   1,
		}),
	}
}

func testMain(m *testing.M) int {
	var cancel context.CancelFunc
	testCTX, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	errCTX, cancel = context.WithCancel(testCTX)
	cancel()

	testAuthStorage = connectRedis()
	defer func() {
		if err := testAuthStorage.Close(); err != nil {
			panic(err)
		}
	}()

	return m.Run()
}

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func TestAuthStorage_Close(t *testing.T) {
	as := connectRedis()
	if err := as.Close(); err != nil {
		t.Fatal(err)
	}
}

func Test_authStorage_Health(t *testing.T) {
	tests := []struct {
		name    string
		ctx     context.Context
		wantErr bool
	}{
		{
			"context error",
			errCTX,
			true,
		},
		{
			"success",
			testCTX,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testAuthStorage.Health(tt.ctx); (err != nil) != tt.wantErr {
				t.Errorf("authStorage.Health() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
