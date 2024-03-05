package datastore

import (
	"context"
	"errors"

	"github.com/op_zero/authserver/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (td *TokenDatastore) AddClient(ctx context.Context, client model.Client) error {
	coll := td.db.Database(USER_DATABASE).Collection(CLIENT_COLLECTION)
	_, err := coll.InsertOne(ctx, client)
	if err != nil {
		return err
	}
	return nil
}

func (td *TokenDatastore) GetClientByClientId(ctx context.Context, clientId string) (*model.Client, error) {
	var client model.Client
	coll := td.db.Database(USER_DATABASE).Collection(CLIENT_COLLECTION)
	filter := bson.M{"clientId": clientId}
	err := coll.FindOne(ctx, filter).Decode(&client)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("no client found")
	} else if err != nil {
		return nil, err
	}
	return &client, err
}
