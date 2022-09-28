package usecase

import (
	"bookhub/internal/entity"
	"context"
	"errors"
	"fmt"
)

type BookStorage interface {
	CreateBook(ctx context.Context, repoId int64, book entity.Book) (int64, error)
	GetBookFile(ctx context.Context, bookID int64) (entity.File, error)
	GetBookImage(ctx context.Context, bookID int64) (entity.File, error)
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

func (bs *BookService) CreateBook(ctx context.Context, repoId int64, book entity.Book) (ID int64, err error) {
	if len(book.Image.File) != 0 {
		book.Image.Name = book.Title + "_" + book.Author
		book.Image.Path, err = bs.fileManager.CreateFile(ctx, book.Image)
		if err != nil && !errors.Is(err, entity.ErrFileAlreadyExists) {
			return 0, fmt.Errorf("CreateFile: %w", err)
		}
	}
	book.File.Name = book.Title + "_" + book.Author
	book.File.Path, err = bs.fileManager.CreateFile(ctx, book.File)
	if err != nil && !errors.Is(err, entity.ErrFileAlreadyExists) {
		return 0, fmt.Errorf("CreateFile: %w", err)
	}

	ID, err = bs.storage.CreateBook(ctx, repoId, book)
	if err != nil {
		return 0, fmt.Errorf("CreateBook: %w", err)
	}
	return ID, nil
}
func (bs *BookService) GetBookFile(ctx context.Context, bookID int64) (file entity.File, err error) {
	book, err := bs.storage.GetBook(ctx, bookID)
	if err != nil {
		return entity.File{}, fmt.Errorf("GetBook: %w", err)
	}
	file, err = bs.fileManager.GetFile(ctx, book.File.Path)
	if err != nil {
		return entity.File{}, fmt.Errorf("GetFile: %w", err)
	}
	return file, nil
}
func (bs *BookService) GetBookImage(ctx context.Context, bookID int64) (file entity.File, err error) {
	book, err := bs.storage.GetBook(ctx, bookID)
	if err != nil {
		return entity.File{}, fmt.Errorf("GetBook: %w", err)
	}
	if book.Image.Path == "" {
		return entity.File{}, entity.ErrImageNotFound
	}
	file, err = bs.fileManager.GetFile(ctx, book.Image.Path)
	if err != nil {
		return entity.File{}, fmt.Errorf("GetFile: %w", err)
	}
	return file, nil
}

func (bs *BookService) GetBook(ctx context.Context, ID int64) (book entity.Book, err error) {
	book, err = bs.storage.GetBook(ctx, ID)
	if err != nil {
		return entity.Book{}, fmt.Errorf("GetBook: %w", err)
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
func (bs *BookService) UpdateBook(ctx context.Context, newBook entity.Book) (book entity.Book, err error) {
	book, err = bs.storage.GetBook(ctx, newBook.ID)
	if err != nil {
		return entity.Book{}, fmt.Errorf("GetBook: %w", err)
	}

	book = bs.updateBook(book, newBook)
	err = bs.storage.UpdateBook(ctx, book)
	if err != nil {
		return entity.Book{}, fmt.Errorf("UpdateBook: %w", err)
	}
	return book, nil
}

func (bs *BookService) DeleteBook(ctx context.Context, ID int64) (err error) {
	book, err := bs.storage.GetBook(ctx, ID)
	if err != nil {
		return fmt.Errorf("GetBook: %w", err)
	}
	if book.Image.Path != "" {
		err = bs.fileManager.DeleteFile(ctx, book.Image.Path)
		if err != nil {
			return fmt.Errorf("DeleteFile: %w", err)
		}
	}
	err = bs.fileManager.DeleteFile(ctx, book.File.Path)
	if err != nil {
		return fmt.Errorf("DeleteFile: %w", err)
	}
	err = bs.storage.DeleteBook(ctx, ID)
	if err != nil {
		return fmt.Errorf("DeleteFile: %w", err)
	}
	return nil
}
