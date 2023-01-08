package storage

import (
	"context"
	"github.com/VeneLooool/BookHub/internal/entity"
)

type Storage interface {
	UserStorage
	RepoStorage
	BookStorage
	FileManager
}

type UserStorage interface {
	CreateUser(context.Context, entity.User) (int64, error)
	GetUser(context.Context, int64) (entity.User, error)
	UpdateUser(context.Context, entity.User) error
	DeleteUser(context.Context, int64) error
}

type RepoStorage interface {
	CreateRepo(ctx context.Context, userID int64, repo entity.Repo) (int64, error)
	GetReposForUser(ctx context.Context, userID int64) ([]entity.Repo, error)
	GetRepo(context.Context, int64) (entity.Repo, error)
	UpdateRepo(context.Context, entity.Repo) error
	DeleteRepo(context.Context, int64) error
	DeleteBookFromRepo(ctx context.Context, RepoID, bookID int64) error
}

type BookStorage interface {
	CreateBook(ctx context.Context, repoId int64, book entity.Book) (int64, error)
	GetBooksForRepo(ctx context.Context, repoID int64) ([]entity.Book, error)
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

type CacheAble interface {
	entity.Book | entity.Repo | entity.User
}

type Cache[Key comparable, T CacheAble] interface {
	Set(Key, T) error
	Get(Key) (T, error)
	Exist(Key) (bool, error)
	Update(Key, T) error
	Delete(Key) error
}
