package service

import (
	"context"

	"github.com/dmuthuraaj/usermgt/model"
)

type Service interface {
	CreateTenant(ctx context.Context, tenant model.Tenant) error
}
