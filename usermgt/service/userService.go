package service

import (
	"context"
	"errors"
	"time"

	"github.com/dmuthuraaj/op_zero/usermgt/datastore"
	"github.com/dmuthuraaj/op_zero/usermgt/model"
	"github.com/dmuthuraaj/op_zero/usermgt/utils"
	"github.com/google/uuid"
)

const (
	automatedProfileUrl = "https://ui-avatars.com/api/?name="
)

var (
	ErrUserAlreadyAdded = errors.New("user already added")
)

type UserService struct {
	userDatastore datastore.DataStore
}

func NewUserService(td datastore.DataStore) *UserService {
	return &UserService{userDatastore: td}
}

func (us *UserService) CreateUser(ctx context.Context, user *model.User) error {
	var err error
	t, err := us.userDatastore.GetUserByName(ctx, user.UserName)
	if err == nil && t != nil {
		return ErrUserAlreadyAdded
	}
	user.Identifier = uuid.New().String()
	user.Password, _ = utils.GenerateFromPassword(user.Password)
	user.Active = true
	user.ProfileUrl = automatedProfileUrl + user.UserName
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err = us.userDatastore.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	users, err := us.userDatastore.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) GetUserByName(ctx context.Context, teantName string) (*model.User, error) {
	var user *model.User
	// TODO: Get UserName from context (stored from the auth middleware)
	user, err := us.userDatastore.GetUserByName(ctx, teantName)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) UpdateUser(ctx context.Context, user *model.User) error {
	var err error
	// TODO: Get UserName from context (stored from the auth middleware)
	u, err := us.userDatastore.GetUserByName(ctx, user.UserName)
	if err != nil {
		return err
	}
	u.UpdatedAt = time.Now()
	err = us.userDatastore.UpdateUser(ctx, u)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) DeleteUserByName(ctx context.Context, teantName string) error {
	var err error
	// TODO: Have to Delete all the info related(associated with) to User
	user, err := us.userDatastore.GetUserByName(ctx, teantName)
	if err != nil {
		return err
	}
	err = us.userDatastore.DeleteUserByName(ctx, user.UserName)
	if err != nil {
		return err
	}
	return nil
}
