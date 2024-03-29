package datastore

import (
	"context"
	"errors"
	"os"

	"github.com/dmuthuraaj/op_zero/tenantmgt/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	LOCAL_URI         = "mongodb://localhost:27017"
	DATABASE          = "tenantadm"
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

func (t *TenantDatastore) CreateTenant(ctx context.Context, tenant *model.Tenant) error {
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

func (t *TenantDatastore) UpdateTenantContactInfo(ctx context.Context, tenant *model.Tenant) error {
	coll := t.db.Database(DATABASE).Collection(TENANT_COLLECTION)
	filter := bson.M{"_id": tenant.Identifier}
	update := bson.M{"$set": bson.M{
		"contactInfo": bson.M{
			"name":         tenant.ContactInfo.Name,
			"email":        tenant.ContactInfo.Email,
			"mobileNumber": tenant.ContactInfo.MobileNumber,
		},
		"updatedAt": tenant.UpdatedAt,
	}}
	res, err := coll.UpdateOne(ctx, filter, update)
	if res.MatchedCount != 1 {
		return ErrTenantNotFound
	} else if err != nil {
		return err
	}
	return nil
}

func (t *TenantDatastore) DeleteTenantByName(ctx context.Context, tenantName string) error {
	coll := t.db.Database(DATABASE).Collection(TENANT_COLLECTION)
	filter := bson.M{"tenantName": tenantName}
	res, err := coll.DeleteOne(ctx, filter)
	if res.DeletedCount != 1 {
		return ErrTenantNotFound
	} else if err != nil {
		return err
	}
	return nil
}
