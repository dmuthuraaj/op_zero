package datastore

import (
	"context"

	"github.com/dmuthuraaj/usermgt/model"
)

type DataStore interface {
	CreateTenant(context.Context, model.Tenant) error
	GetTenantByName(ctx context.Context, tenantName string) (*model.Tenant, error)
}