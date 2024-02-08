package service

import (
	"context"

	"github.com/dmuthuraaj/op_zero/usermgt/model"
)

type Service interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetAllUsers(ctx context.Context) ([]model.User, error)
	GetUserByName(ctx context.Context, teantName string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUserByName(ctx context.Context, teantName string) error
}
