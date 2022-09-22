package storage

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

type RepoStorage interface {
	CreateRepo(context.Context, entity.Repo) (int64, error)
	GetRepo(context.Context, int64) (entity.Repo, error)
	UpdateRepo(context.Context, entity.Repo) error
	DeleteRepo(context.Context, int64) error
}

type BookStorage interface {
	CreateBook(context.Context, entity.Book) (int64, error)
	GetBook(context.Context, int64) (entity.Book, error)
	UpdateBook(context.Context, entity.Book) error
	DeleteBook(context.Context, int64) error
}

type FileManager interface {
	CreateFile(context.Context, entity.File) (string, error)
	GetFile(context.Context, string) (entity.File, error)
	UpdateFile(context.Context, entity.File) error
	DeleteFile(context.Context, string) error
}

type Storage interface {
	UserStorage
	RepoStorage
	BookStorage
	FileManager
}
