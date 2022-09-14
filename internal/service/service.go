package service

import (
	"bookhub/internal/entity"
	"bookhub/internal/storage"
	"context"
)

type BookUseCase interface {
	CreateBook(context.Context, entity.Book) (int64, error)
	GetBook(context.Context, int64) (entity.Book, error)
	UpdateBook(context.Context, entity.Book) error
	DeleteBook(context.Context, int64) error
}
type UserUseCase interface {
	CreateUser(context.Context, entity.User) (int64, error)
	GetUser(context.Context, int64) (entity.User, error)
	UpdateUser(context.Context, entity.User) error
	DeleteUser(context.Context, int64) error
}
type RepoUseCase interface {
	CreateRepo(context.Context, entity.Repo) (int64, error)
	GetRepo(context.Context, int64) (entity.Repo, error)
	UpdateRepo(context.Context, entity.Repo) error
	DeleteRepo(context.Context, int64) error
}

type Service struct {
	User UserUseCase
	Book BookUseCase
	Repo RepoUseCase
}

type Deps struct {
	Storage *storage.Storage
}

/*
func NewService(deps Deps) *Service {
	userService := usecase.NewUserService(deps.Storage.UserStorage)
	bookService := usecase.NewBookService(deps.Storage.BookStorage)
	repoService := usecase.NewRepoService(deps.Storage.RepoStorage)
	return &Service{
		User: userService,
		Book: bookService,
		Repo: repoService,
	}
}
*/
//TODO как лучше это сделать?

func NewService(userService UserUseCase, bookService BookUseCase, repoService RepoUseCase) *Service {
	return &Service{
		User: userService,
		Book: bookService,
		Repo: repoService,
	}
}
