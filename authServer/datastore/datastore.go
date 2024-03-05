package datastore

import (
	"context"

	"github.com/op_zero/authserver/model"
)

type Datastore interface {
	TDatastore
	UDatastore
}

type TDatastore interface {
	SaveAccessToken(ctx context.Context, accessToken *model.OauthAccessTokenResponse) error
	SaveRefreshToken(ctx context.Context, refreshToken *model.RefreshToken) error
}

type UDatastore interface {
	AddClient(ctx context.Context, client model.Client) error
	GetClientByClientId(ctx context.Context, clientId string) (*model.Client, error)
}
