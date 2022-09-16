package service

import (
	"bookhub/internal/entity"
	"context"
)

type BookUseCase interface {
	CreateBook(context.Context, entity.Book) (int64, error)
	GetBook(context.Context, int64) (entity.Book, error)
	UpdateBook(context.Context, entity.Book) (entity.Book, error)
	DeleteBook(context.Context, int64) error
}
type UserUseCase interface {
	CreateUser(context.Context, entity.User) (int64, error)
	GetUser(context.Context, int64) (entity.User, error)
	UpdateUser(context.Context, entity.User) (entity.User, error)
	DeleteUser(context.Context, int64) error
}
type RepoUseCase interface {
	CreateRepo(context.Context, entity.Repo) (int64, error)
	GetRepo(context.Context, int64) (entity.Repo, error)
	UpdateRepo(context.Context, entity.Repo) (entity.Repo, error)
	DeleteRepo(context.Context, int64) error
}

type Service interface {
	UserUseCase
	BookUseCase
	RepoUseCase
}
