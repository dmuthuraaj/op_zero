package datastore

import (
	"context"

	"github.com/dmuthuraaj/op_zero/usermgt/model"
)

type DataStore interface {
	CreateUser(context.Context, *model.User) error
	GetAllUsers(ctx context.Context) ([]model.User, error)
	GetUserByName(ctx context.Context, userName string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUserByName(ctx context.Context, userName string) error
}
