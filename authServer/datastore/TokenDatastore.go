package datastore

import (
	"context"

	"github.com/op_zero/authserver/config"
	"github.com/op_zero/authserver/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	MONGO_LOCAL_URI          = "mongodb://localhost:27017"
	MONGO_LOCAL_USERNAME     = "admin"
	MONGO_LOCAL_PASSWORD     = "pass"
	TOKEN_DATABASE           = "token"
	USER_DATABASE            = "user"
	CLIENT_COLLECTION        = "client"
	ACCESS_TOKEN_COLLECTION  = "access_token"
	REFRESH_TOKEN_COLLECTION = "refresh_token"
	AUTH_CODE_COLLECTION     = "authorization_code"
)

type TokenDatastore struct {
	db *mongo.Client
}

func NewTokenDatastore(config config.Config) (*TokenDatastore, error) {
	var err error
	var uri = config.Mongo.Uri
	if uri == "" {
		uri = MONGO_LOCAL_URI
	}
	opts := options.Client()
	opts.ApplyURI(uri)
	if username := config.Mongo.UserName; username == "" {
		username = MONGO_LOCAL_USERNAME
	}
	if password := config.Mongo.Password; password == "" {
		password = MONGO_LOCAL_PASSWORD
	}
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), &readpref.ReadPref{})
	if err != nil {
		return nil, err
	}
	return &TokenDatastore{db: client}, nil
}

func (tb *TokenDatastore) SaveAccessToken(ctx context.Context, accessToken *model.OauthAccessTokenResponse) error {
	coll := tb.db.Database(TOKEN_DATABASE).Collection(ACCESS_TOKEN_COLLECTION)
	_, err := coll.InsertOne(ctx, accessToken)
	if err != nil {
		return err
	}
	return nil
}

func (tb *TokenDatastore) SaveRefreshToken(ctx context.Context, refreshToken *model.RefreshToken) error {
	coll := tb.db.Database(TOKEN_DATABASE).Collection(REFRESH_TOKEN_COLLECTION)
	_, err := coll.InsertOne(ctx, refreshToken)
	if err != nil {
		return err
	}
	return nil
}
