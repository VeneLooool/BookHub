package postgres

import (
	"bookhub/internal/entity"
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type BookStorage struct {
	db *sqlx.DB
}

func (st *BookStorage) CreateBook(ctx context.Context, repoId int64, book entity.Book) (ID int64, err error) {
	tx, err := st.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("BeginTxx: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("transcation failed, recover from panic %w", err)
			if errRB := tx.Rollback(); errRB != nil {
				err = fmt.Errorf("%w rollback failed: %w", err, errRB)
			}
		} else if err != nil {
			err = fmt.Errorf("transaction failed, rollback: %w", err)
			if errRB := tx.Rollback(); errRB != nil {
				err = fmt.Errorf("%w rollback failed: %w", err, errRB)
			}
		} else {
			if err = tx.Commit(); err != nil {
				err = fmt.Errorf("transaction commit failed: %w", err)
			}
		}

	}()

	var b entity.Book
	if err = tx.QueryRowxContext(ctx, createBook,
		&book.Title,
		&book.Author,
		&book.NumberPages,
		&book.Desc,
		&book.Image.Path,
		&book.File.Path).Scan(b); err != nil {
		return 0, fmt.Errorf("QueryRowxContext: %w", err)
	}
	result, err := tx.ExecContext(ctx, attachBookToRepo, &b.ID, &b.CurrentPage, &repoId)
	if err != nil {
		return 0, fmt.Errorf("ExecContext: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("RowsAffected: %w", err)
	} else if rowsAffected == 0 {
		return 0, errors.New("attachBookToRepo: the lines were not affected")
	}
	return b.ID, nil
}
func (st *BookStorage) GetBook(ctx context.Context, ID int64) (book entity.Book, err error) {
	if err = st.db.GetContext(ctx, &book, getBook, &ID); err != nil {
		return entity.Book{}, fmt.Errorf("GetContext: %w", err)
	}
	return book, nil
}
func (st *BookStorage) UpdateBook(ctx context.Context, book entity.Book) error {
	var b entity.Book
	if err := st.db.QueryRowxContext(ctx, updateBook,
		&book.Title,
		&book.Author,
		&book.NumberPages,
		&book.Desc,
		&book.Image.Path,
		&book.File.Path,
		&book.ID).Scan(&b); err != nil {
		return fmt.Errorf("QueryRowxContext: %w", err)
	}
	return nil
}
func (st *BookStorage) DeleteBook(ctx context.Context, ID int64) error {
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
	return nil
}
