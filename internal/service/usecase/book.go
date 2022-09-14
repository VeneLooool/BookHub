package usecase

import (
	"bookhub/internal/entity"
	"context"
)

type BookStorage interface {
	CreateBook(context.Context, entity.Book) (int64, error)
	GetBook(context.Context, int64) (entity.Book, error)
	UpdateBook(context.Context, entity.Book) error
	DeleteBook(context.Context, int64) error
}

type BookService struct {
	storage BookStorage
}

func NewBookService(bookStorage BookStorage) *BookService {
	return &BookService{
		storage: bookStorage,
	}
}
