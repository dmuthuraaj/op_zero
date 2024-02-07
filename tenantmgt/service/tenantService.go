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

func (ts *TenantService) CreateTenant(ctx context.Context, tenant *model.Tenant) error {
	var err error
	t, err := ts.tenantDatastore.GetTenantByName(ctx, tenant.TenantName)
	if err == nil && t != nil {
		return ErrTenantAlreadyAdded
	}
	tenant.Identifier = uuid.New().String()
	tenant.State = model.STATE_ACTIVE
	tenant.CreatedAt = time.Now()
	tenant.UpdatedAt = time.Now()
	err = ts.tenantDatastore.CreateTenant(ctx, tenant)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TenantService) GetTenantByName(ctx context.Context, teantName string) (*model.Tenant, error) {
	var tenant *model.Tenant
	// TODO: Get TenantName from context (stored from the auth middleware)
	tenant, err := ts.tenantDatastore.GetTenantByName(ctx, teantName)
	if err != nil {
		return nil, err
	}
	return tenant, nil
}

func (ts *TenantService) UpdateTenantContactInfo(ctx context.Context, tenant *model.Tenant) error {
	var err error
	// TODO: Get TenantName from context (stored from the auth middleware)
	t, err := ts.tenantDatastore.GetTenantByName(ctx, tenant.TenantName)
	if err != nil {
		return err
	}
	t.UpdatedAt = time.Now()
	if t.ContactInfo.Name != tenant.ContactInfo.Name {
		t.ContactInfo.Name = tenant.ContactInfo.Name
	}
	if t.ContactInfo.Email != tenant.ContactInfo.Email {
		t.ContactInfo.Email = tenant.ContactInfo.Email
	}
	if t.ContactInfo.MobileNumber != tenant.ContactInfo.MobileNumber {
		t.ContactInfo.MobileNumber = tenant.ContactInfo.MobileNumber
	}
	err = ts.tenantDatastore.UpdateTenantContactInfo(ctx, t)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TenantService) DeleteTenantByName(ctx context.Context, teantName string) error {
	var err error
	// TODO: Have to Delete all the info related(associated with) to Tenant
	tenant, err := ts.tenantDatastore.GetTenantByName(ctx, teantName)
	if err != nil {
		return err
	}
	err = ts.tenantDatastore.DeleteTenantByName(ctx, tenant.TenantName)
	if err != nil {
		return err
	}
	return nil
}
