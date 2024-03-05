package services

import (
	"context"

	"github.com/op_zero/authserver/model"
)

type Service interface {
	CreateAccessTokenForClientCredentials(ctx context.Context, tokenReq *model.OauthTokenRequest) (*model.OauthAccessTokenResponse, error)
}
