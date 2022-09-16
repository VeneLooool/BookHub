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

func (bs *BookService) CreateBook(ctx context.Context, book entity.Book) (int64, error) {
	return bs.storage.CreateBook(ctx, book)
}
func (bs *BookService) GetBook(ctx context.Context, ID int64) (entity.Book, error) {
	return bs.storage.GetBook(ctx, ID)
}
func (bs *BookService) updateBook(oldBook, newBook entity.Book) entity.Book {
	if newBook.Title != "" {
		oldBook.Title = newBook.Title
	}
	if newBook.Desc != "" {
		oldBook.Desc = newBook.Desc
	}
	if newBook.Author != "" {
		oldBook.Author = newBook.Author
	}
	if newBook.NumberPages != 0 {
		oldBook.NumberPages = newBook.NumberPages
	}
	return oldBook
}
func (bs *BookService) UpdateBook(ctx context.Context, newBook entity.Book) (entity.Book, error) {
	book, err := bs.GetBook(ctx, newBook.ID)
	if err != nil {
		return entity.Book{}, err
	}

	book = bs.updateBook(book, newBook)
	err = bs.storage.UpdateBook(ctx, book)
	if err != nil {
		return entity.Book{}, err
	}
	return book, err
}

func (bs *BookService) DeleteBook(ctx context.Context, ID int64) error {
	return bs.storage.DeleteBook(ctx, ID)
}
