package oidcredis

import (
	"time"

	"github.com/muhlemmer/oidc-redis/internal/model"
	"github.com/zitadel/oidc/pkg/oidc"
	"github.com/zitadel/oidc/pkg/op"
)

// Request implements both AuthRequest and RefreshTokenRequest,
// as they have overlapping properties.
// Unused fieldnames are automatically hidden by the interface
// types uses in AuthStorage.
type Request struct {
	model.Request
}

func (req *Request) GetCodeChallenge() *oidc.CodeChallenge {
	cc := req.Request.GetCodeChallenge()
	return &oidc.CodeChallenge{
		Challenge: cc.GetChallenge(),
		Method:    oidc.CodeChallengeMethod(cc.GetMethod().String()),
	}
}

func (req *Request) GetAuthTime() time.Time {
	return req.Request.GetAuthTime().AsTime()
}

func (req *Request) GetResponseType() oidc.ResponseType {
	return oidc.ResponseType(req.Request.ResponseType)
}

func (req *Request) GetResponseMode() oidc.ResponseMode {
	return oidc.ResponseMode(req.Request.GetResponseMode())
}

func (req *Request) Done() bool {
	return req.GetDone()
}

func (req *Request) SetCurrentScopes(scopes []string) {
	req.Scopes = scopes
}

// implementation check
func _() op.AuthRequest         { return new(Request) }
func _() op.RefreshTokenRequest { return new(Request) }
