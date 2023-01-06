package storage

import (
	"context"
	"fmt"
	"github.com/VeneLooool/BookHub/internal/entity"
	"strconv"
)

type BookStorageAbs struct {
	db    BookStorage
	cache Cache[string, entity.Book]
}

func NewBookStorageAbs(db BookStorage, ch Cache[string, entity.Book]) *BookStorageAbs {
	return &BookStorageAbs{
		db:    db,
		cache: ch,
	}
}

func (bsa *BookStorageAbs) CreateBook(ctx context.Context, repoId int64, book entity.Book) (ID int64, err error) {
	ID, err = bsa.db.CreateBook(ctx, repoId, book)
	if err != nil {
		return 0, fmt.Errorf("db CreateBook: %w", err)
	}

	key := bsa.idToCacheKey(ID)
	if err = bsa.cache.Set(key, book); err != nil {
		//TODO add logger
	}
	return ID, nil
}
func (bsa *BookStorageAbs) GetBooksForRepo(ctx context.Context, repoID int64) ([]entity.Book, error) {
	return bsa.db.GetBooksForRepo(ctx, repoID)
}
func (bsa *BookStorageAbs) GetBook(ctx context.Context, ID int64) (book entity.Book, err error) {
	key := bsa.idToCacheKey(ID)
	exist, err := bsa.cache.Exist(key)
	if err != nil {
		//TODO add logger
	}
	if exist {
		book, err = bsa.cache.Get(key)
		if err == nil {
			return book, nil
		}
		//TODO add logger
	}

	book, err = bsa.db.GetBook(ctx, ID)
	if err != nil {
		return entity.Book{}, fmt.Errorf("db GetBook: %w", err)
	}
	return book, nil
}
func (bsa *BookStorageAbs) UpdateBook(ctx context.Context, book entity.Book) error {
	err := bsa.db.UpdateBook(ctx, book)
	if err != nil {
		return fmt.Errorf("db UpdateBook: %w", err)
	}

	key := bsa.idToCacheKey(book.ID)
	if err = bsa.cache.Update(key, book); err != nil {
		//TODO add logger
	}
	return nil
}
func (bsa *BookStorageAbs) DeleteBook(ctx context.Context, ID int64) error {
	err := bsa.db.DeleteBook(ctx, ID)
	if err != nil {
		return fmt.Errorf("db DeleteBook: %w", err)
	}

	key := bsa.idToCacheKey(ID)
	if err = bsa.cache.Delete(key); err != nil {
		//TODO add logger
	}
	return nil
}
func (bsa *BookStorageAbs) idToCacheKey(ID int64) string {
	return strconv.Itoa(int(ID))
}
