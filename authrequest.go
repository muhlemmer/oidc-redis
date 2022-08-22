package oidcredis

import (
	"time"

	"github.com/muhlemmer/oidc-redis/internal/model"
	"github.com/zitadel/oidc/pkg/oidc"
	"github.com/zitadel/oidc/pkg/op"
)

type AuthRequest struct {
	model.AuthRequest
}

func (req *AuthRequest) GetCodeChallenge() *oidc.CodeChallenge {
	cc := req.AuthRequest.GetCodeChallenge()
	return &oidc.CodeChallenge{
		Challenge: cc.GetChallenge(),
		Method:    oidc.CodeChallengeMethod(cc.GetMethod().String()),
	}
}

func (req *AuthRequest) GetAuthTime() time.Time {
	return req.AuthRequest.GetAuthTime().AsTime()
}

func (req *AuthRequest) GetResponseType() oidc.ResponseType {
	return oidc.ResponseType(req.AuthRequest.ResponseType)
}

func (req *AuthRequest) GetResponseMode() oidc.ResponseMode {
	return oidc.ResponseMode(req.AuthRequest.GetResponseMode())
}

func (req *AuthRequest) Done() bool {
	return req.GetDone()
}

// implementation check
func _() op.AuthRequest {
	return new(AuthRequest)
}
