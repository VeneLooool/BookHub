package usecase

import (
	"bookhub/internal/entity"
	"context"
	"errors"
	"fmt"
)

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

type BookService struct {
	storage     BookStorage
	fileManager FileManager
}

func NewBookService(bookStorage BookStorage, fileManager FileManager) *BookService {
	return &BookService{
		storage:     bookStorage,
		fileManager: fileManager,
	}
}

func (bs *BookService) CreateBook(ctx context.Context, book entity.Book) (int64, error) {
	var err error

	if len(book.Image.File) != 0 {
		book.Image.Path, err = bs.fileManager.CreateFile(ctx, book.Image)
		if err != nil {
			return 0, fmt.Errorf("CreateBook: unable to create image: %w", err)
		}
	}
	if len(book.File.File) == 0 {
		return 0, errors.New("CreateBook: book file not provided")
	}
	book.File.Path, err = bs.fileManager.CreateFile(ctx, book.File)
	if err != nil {
		return 0, fmt.Errorf("CreateBook: unable to create book file: %w", err)
	}

	return bs.storage.CreateBook(ctx, book)
}
func (bs *BookService) GetBook(ctx context.Context, ID int64) (entity.Book, error) {
	book, err := bs.storage.GetBook(ctx, ID)
	if err != nil {
		return entity.Book{}, fmt.Errorf("GetBook: unable to get book: %w", err)
	}
	if book.Image.Path != "" {
		book.Image, err = bs.fileManager.GetFile(ctx, book.Image.Path)
		if err != nil {
			return entity.Book{}, fmt.Errorf("GetBook: unable to get image: %w", err)
		}
	}
	if book.File.Path == "" {
		return entity.Book{}, errors.New("GetBook: empty book path")
	}
	book.File, err = bs.fileManager.GetFile(ctx, book.File.Path)
	if err != nil {
		return entity.Book{}, fmt.Errorf("GetBook: unable to get book file: %w", err)
	}
	return book, nil
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
	book, err := bs.storage.GetBook(ctx, newBook.ID)
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
	book, err := bs.storage.GetBook(ctx, ID)
	if err != nil {
		return fmt.Errorf("DeleteBook: unable to get book: %w", err)
	}
	if book.Image.Path != "" {
		err = bs.fileManager.DeleteFile(ctx, book.Image.Path)
		if err != nil {
			return fmt.Errorf("DeleteBook: unable to delete image: %w", err)
		}
	}
	err = bs.fileManager.DeleteFile(ctx, book.File.Path)
	if err != nil {
		return fmt.Errorf("DeleteBook: unable to delete book file: %w", err)
	}
	return bs.storage.DeleteBook(ctx, ID)
}
