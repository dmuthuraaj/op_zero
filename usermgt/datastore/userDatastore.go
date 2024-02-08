package datastore

import (
	"context"
	"errors"
	"os"

	"github.com/dmuthuraaj/op_zero/usermgt/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	LOCAL_URI       = "mongodb://localhost:27017"
	DATABASE        = "useradm"
	USER_COLLECTION = "users"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserDatastore struct {
	db *mongo.Client
}

func NewUserDatastore() *UserDatastore {
	var err error
	var client *mongo.Client
	opts := options.Client()
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = LOCAL_URI
	}
	opts.ApplyURI(uri)
	client, err = mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), &readpref.ReadPref{})
	if err != nil {
		panic(err)
	}
	return &UserDatastore{db: client}
}

func (u *UserDatastore) CreateUser(ctx context.Context, user *model.User) error {
	coll := u.db.Database(DATABASE).Collection(USER_COLLECTION)
	_, err := coll.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserDatastore) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	filter := bson.M{}
	var err error
	coll := u.db.Database(DATABASE).Collection(USER_COLLECTION)
	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	err = cur.Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserDatastore) GetUserByName(ctx context.Context, userName string) (*model.User, error) {
	var user model.User
	coll := u.db.Database(DATABASE).Collection(USER_COLLECTION)
	filter := bson.M{"userName": userName}
	err := coll.FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDatastore) UpdateUser(ctx context.Context, user *model.User) error {
	coll := u.db.Database(DATABASE).Collection(USER_COLLECTION)
	filter := bson.M{"_id": user.Identifier}
	update := bson.M{"$set": bson.M{
		"updatedAt": user.UpdatedAt,
	}}
	res, err := coll.UpdateOne(ctx, filter, update)
	if res.MatchedCount != 1 {
		return ErrUserNotFound
	} else if err != nil {
		return err
	}
	return nil
}

func (u *UserDatastore) DeleteUserByName(ctx context.Context, userName string) error {
	coll := u.db.Database(DATABASE).Collection(USER_COLLECTION)
	filter := bson.M{"userName": userName}
	res, err := coll.DeleteOne(ctx, filter)
	if res.DeletedCount != 1 {
		return ErrUserNotFound
	} else if err != nil {
		return err
	}
	return nil
}
