package storage

import (
	"bookhub/internal/entity"
	"context"
)

type UserStorage interface {
	CreateUser(context.Context, entity.User) (int64, error)
	GetUser(context.Context, int64) (entity.User, error)
	GetReposForUser(ctx context.Context, userID int64) ([]entity.Repo, error)
	UpdateUser(context.Context, entity.User) error
	DeleteUser(context.Context, int64) error
}

type RepoStorage interface {
	CreateRepo(ctx context.Context, userID int64, repo entity.Repo) (int64, error)
	GetRepo(context.Context, int64) (entity.Repo, error)
	GetBooksForRepo(ctx context.Context, repoID int64) ([]entity.Book, error)
	UpdateRepo(context.Context, entity.Repo) error
	DeleteRepo(context.Context, int64) error
	DeleteBookFromRepo(ctx context.Context, RepoID, bookID int64) error
}

type BookStorage interface {
	CreateBook(ctx context.Context, repoId int64, book entity.Book) (int64, error)
	GetBook(context.Context, int64) (entity.Book, error)
	UpdateBook(context.Context, entity.Book) error
	DeleteBook(context.Context, int64) error
}

type FileManager interface {
	CreateFile(ctx context.Context, file entity.File) (path string, err error)
	GetFile(ctx context.Context, path string) (file entity.File, err error)
	UpdateFile(ctx context.Context, file entity.File) error
	DeleteFile(ctx context.Context, path string) error
}

type Storage interface {
	UserStorage
	RepoStorage
	BookStorage
	FileManager
}
