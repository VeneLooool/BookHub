package service

import (
	"bookhub/internal/entity"
	"context"
)

type BookUseCase interface {
	CreateBook(ctx context.Context, repoId int64, book entity.Book) (int64, error)
	GetBook(ctx context.Context, ID int64) (entity.Book, error)
	GetBookFile(ctx context.Context, bookID int64) (entity.File, error)
	GetBookImage(ctx context.Context, bookID int64) (entity.File, error)
	GetBooksForRepo(ctx context.Context, repoID int64) ([]entity.Book, error)
	UpdateBook(ctx context.Context, book entity.Book) (entity.Book, error)
	DeleteBook(ctx context.Context, ID int64) error
}
type UserUseCase interface {
	CreateUser(ctx context.Context, user entity.User) (int64, error)
	GetUser(ctx context.Context, ID int64) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, ID int64) error
}
type RepoUseCase interface {
	CreateRepo(ctx context.Context, userID int64, repo entity.Repo) (int64, error)
	GetReposForUser(ctx context.Context, userID int64) ([]entity.Repo, error)
	GetRepo(ctx context.Context, ID int64) (entity.Repo, error)
	UpdateRepo(ctx context.Context, repo entity.Repo) (entity.Repo, error)
	DeleteRepo(ctx context.Context, ID int64) error
	DeleteBookFromRepo(ctx context.Context, RepoID, bookID int64) error
}

type Service interface {
	UserUseCase
	BookUseCase
	RepoUseCase
}
