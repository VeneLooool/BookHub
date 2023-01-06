package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/VeneLooool/BookHub/internal/entity"
	"github.com/VeneLooool/BookHub/internal/storage"
	"github.com/jmoiron/sqlx"
)

type BookStorage struct {
	db *sqlx.DB
}

func NewBookStorage(db *sqlx.DB) storage.BookStorage {
	return &BookStorage{db: db}
}

func (st *BookStorage) CreateBook(ctx context.Context, repoId int64, book entity.Book) (ID int64, err error) {
	tx, err := st.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("Begin: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("transcation failed, recover from panic %w", err)
			tx.Rollback()
		} else if err != nil {
			err = fmt.Errorf("transaction failed, rollback: %w", err)
			tx.Rollback()
		} else {
			if err = tx.Commit(); err != nil {
				err = fmt.Errorf("transaction commit failed: %w", err)
			}
		}

	}()

	result, err := tx.ExecContext(ctx, createBook, book.Title, book.Author, book.NumberPages, book.Desc, book.Image.Path, book.File.Path)
	if err != nil {
		return 0, fmt.Errorf("ExecContext: %w", err)
	}
	if book.ID, err = result.LastInsertId(); err != nil {
		return 0, fmt.Errorf("LastInsetrId: %w", err)
	}

	result, err = tx.ExecContext(ctx, attachBookToRepo, book.ID, book.CurrentPage, repoId)
	if err != nil {
		return 0, fmt.Errorf("ExecContext: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("RowsAffected: %w", err)
	}
	if rowsAffected == 0 {
		return 0, errors.New("attachBookToRepo: the lines were not affected")
	}
	return book.ID, nil
}
func (st *BookStorage) GetBook(ctx context.Context, ID int64) (book entity.Book, err error) {
	if err = st.db.GetContext(ctx, &book, getBook, &ID); err != nil {
		return entity.Book{}, fmt.Errorf("GetContext: %w", err)
	}
	return book, nil
}

func (st *BookStorage) GetBooksForRepo(ctx context.Context, repoID int64) (books []entity.Book, err error) {
	if err = st.db.SelectContext(ctx, &books, getBooksForRepo, repoID); err != nil {
		return nil, fmt.Errorf("SelectContext: %w", err)
	}
	return books, nil
}

func (st *BookStorage) UpdateBook(ctx context.Context, book entity.Book) error {
	result, err := st.db.ExecContext(ctx, updateBook, &book.Title, &book.Author, &book.NumberPages,
		&book.Desc, &book.Image.Path, &book.File.Path, &book.ID)
	if err != nil {
		return fmt.Errorf("QueryRowxContext: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("RowsAffected: %w", err)
	}
	if rowsAffected != 1 {
		return errors.New("RowsAffected more or less than one")
	}
	return nil
}
func (st *BookStorage) DeleteBook(ctx context.Context, ID int64) error {
	tx, err := st.db.Begin()
	if err != nil {
		return fmt.Errorf("Begin trransaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("transcation failed, recover from panic %w", err)
			tx.Rollback()
		} else if err != nil {
			err = fmt.Errorf("transaction failed, rollback: %w", err)
			tx.Rollback()
		} else {
			if err = tx.Commit(); err != nil {
				err = fmt.Errorf("transaction commit failed: %w", err)
			}
		}
	}()

	result, err := st.db.ExecContext(ctx, deleteBook, ID)
	if err != nil {
		return fmt.Errorf("ExecContext: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("RowsAffected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("RowsAffceted: %w", entity.ErrBookNotFound)
	}

	if _, err = st.db.ExecContext(ctx, deleteBookFromAllRepos, ID); err != nil {
		return fmt.Errorf("RowsAffected: %w", err)
	}
	return nil
}
