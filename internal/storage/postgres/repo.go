package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/VeneLooool/BookHub/internal/entity"
	"github.com/VeneLooool/BookHub/internal/storage"
	"github.com/jmoiron/sqlx"
)

type RepoStorage struct {
	db *sqlx.DB
}

func NewRepoStorage(db *sqlx.DB) storage.RepoStorage {
	return &RepoStorage{db: db}
}

func (st *RepoStorage) CreateRepo(ctx context.Context, userID int64, repo entity.Repo) (ID int64, err error) {
	err = st.db.QueryRowContext(ctx, createRepo, repo.Name, repo.Visibility, repo.Desc, userID).Scan(&ID)
	if err != nil {
		return 0, fmt.Errorf("QueryRowxContext: %w", err)
	}

	return ID, nil
}
func (st *RepoStorage) GetRepo(ctx context.Context, repoId int64) (repo entity.Repo, err error) {
	if err = st.db.GetContext(ctx, &repo, getRepo, &repoId); err != nil {
		return entity.Repo{}, fmt.Errorf("GetContext: %w", err)
	}
	return repo, nil
}
func (st *RepoStorage) GetReposForUser(ctx context.Context, userID int64) (repos []entity.Repo, err error) {
	if err = st.db.SelectContext(ctx, &repos, getReposForUser, userID); err != nil {
		return nil, fmt.Errorf("SelectContext: %w", err)
	}
	return repos, nil
}

func (st *RepoStorage) UpdateRepo(ctx context.Context, repo entity.Repo) error {
	result, err := st.db.ExecContext(ctx, updateRepo, &repo.Name, &repo.Desc, &repo.ID)
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

func (st *RepoStorage) DeleteRepo(ctx context.Context, repoID int64) error {
	tx, err := st.db.Begin()
	if err != nil {
		return fmt.Errorf("transaction begin: %w", err)
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
	if _, err = st.db.ExecContext(ctx, removeAttachmentsOfBooksToRepo, repoID); err != nil {
		return fmt.Errorf("ExecContext: %w", err)
	}

	result, err := st.db.ExecContext(ctx, deleteRepo, repoID)
	if err != nil {
		return fmt.Errorf("ExecContext: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("RowsAffected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("RowsAffected: %w", entity.ErrRepoNotFound)
	}

	return nil
}

func (st *RepoStorage) DeleteBookFromRepo(ctx context.Context, RepoID, bookID int64) error {
	result, err := st.db.ExecContext(ctx, deleteBookFromRepo, bookID, RepoID)
	if err != nil {
		return fmt.Errorf("ExecContext: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("RowsAffected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("RowsAffected: %w", entity.ErrBookNotFound)
	}
	return nil
}
