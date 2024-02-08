package service

import (
	"context"

	"github.com/dmuthuraaj/op_zero/tenantmgt/model"
)

type Service interface {
	CreateTenant(ctx context.Context, tenant *model.Tenant) error
	GetTenantByName(ctx context.Context, teantName string) (*model.Tenant, error)
	UpdateTenantContactInfo(ctx context.Context, tenant *model.Tenant) error
	DeleteTenantByName(ctx context.Context, teantName string) error
}
