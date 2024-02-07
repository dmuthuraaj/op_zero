package datastore

import (
	"context"
	"errors"
	"os"

	"github.com/dmuthuraaj/usermgt/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	LOCAL_URI         = "mongodb://localhost:27017"
	DATABASE          = "tenant"
	TENANT_COLLECTION = "tenants"
)

var (
	ErrTenantNotFound = errors.New("tenant not found")
)

type TenantDatastore struct {
	db *mongo.Client
}

func NewTenantDatastore() *TenantDatastore {
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
	return &TenantDatastore{db: client}
}

func (t *TenantDatastore) CreateTenant(ctx context.Context, tenant model.Tenant) error {
	coll := t.db.Database(DATABASE).Collection(TENANT_COLLECTION)
	_, err := coll.InsertOne(ctx, tenant)
	if err != nil {
		return err
	}
	return nil
}

func (t *TenantDatastore) GetTenantByName(ctx context.Context, tenantName string) (*model.Tenant, error) {
	var tenant model.Tenant
	coll := t.db.Database(DATABASE).Collection(TENANT_COLLECTION)
	filter := bson.M{"tenantName": tenantName}
	err := coll.FindOne(ctx, filter).Decode(&tenant)
	if err == mongo.ErrNoDocuments {
		return nil, ErrTenantNotFound
	} else if err != nil {
		return nil, err
	}
	return &tenant, nil
}
