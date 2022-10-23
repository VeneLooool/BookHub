package usecase

import (
	"bookhub/internal/entity"
	"bookhub/internal/storage"
	"context"
	"errors"
	"fmt"
)

type BookService struct {
	storage     storage.BookStorage
	fileManager storage.FileManager
}

func NewBookService(bookStorage storage.BookStorage, fileManager storage.FileManager) *BookService {
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
