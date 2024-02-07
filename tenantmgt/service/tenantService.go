package service

import (
	"context"
	"errors"
	"time"

	"github.com/dmuthuraaj/usermgt/datastore"
	"github.com/dmuthuraaj/usermgt/model"
	"github.com/google/uuid"
)

var (
	ErrTenantAlreadyAdded = errors.New("tenant already added")
)

type TenantService struct {
	tenantDatastore datastore.DataStore
}

func NewTenantService(td datastore.DataStore) *TenantService {
	return &TenantService{tenantDatastore: td}
}

func (ts *TenantService) CreateTenant(ctx context.Context, tenant model.Tenant) error {
	var err error
	t, err := ts.tenantDatastore.GetTenantByName(ctx, tenant.TenantName)
	if err == nil && t != nil {
		return ErrTenantAlreadyAdded
	}
	tenant.Identifier = uuid.New().String()
	tenant.State = model.STATE_ACTIVE
	tenant.CreatedAt = time.Now()
	tenant.LastModifiedAt = time.Now()
	err = ts.tenantDatastore.CreateTenant(ctx, tenant)
	if err != nil {
		return err
	}
	return nil
}
