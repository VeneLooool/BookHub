package usecase

import (
	"bookhub/internal/entity"
	"context"
)

type UserStorage interface {
	CreateUser(context.Context, entity.User) (int64, error)
	GetUser(context.Context, int64) (entity.User, error)
	UpdateUser(context.Context, entity.User) error
	DeleteUser(context.Context, int64) error
}

type UserService struct {
	storage UserStorage
}

func NewUserService(userStorage UserStorage) *UserService {
	return &UserService{
		storage: userStorage,
	}
}
