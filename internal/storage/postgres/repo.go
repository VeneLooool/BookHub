package postgres

import (
	"bookhub/internal/entity"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type RepoStorage struct {
	db *sqlx.DB
}

func (st *RepoStorage) CreateRepo(ctx context.Context, userID int64, repo entity.Repo) (ID int64, err error) {
	var r entity.Repo
	if err = st.db.QueryRowxContext(ctx,
		createRepo,
		&repo.Name,
		&repo.Visibility,
		&repo.Desc,
		&userID).Scan(&r); err != nil {
		return 0, fmt.Errorf("QueryRowxContext: %w", err)
	}
	return r.ID, nil
}
func (st *RepoStorage) GetRepo(ctx context.Context, repoId int64) (repo entity.Repo, err error) {
	if err = st.db.GetContext(ctx, &repo, getRepo, &repoId); err != nil {
		return entity.Repo{}, fmt.Errorf("GetContext: %w", err)
	}
	return repo, nil
}
func (st *RepoStorage) GetBooksForRepo(ctx context.Context, repoID int64) (books []entity.Book, err error) {
	rows, err := st.db.QueryxContext(ctx, getBooksForRepo, &repoID)
	if err != nil {
		return nil, fmt.Errorf("QueryxContext: %w", err)
	}
	for rows.Next() {
		var book entity.Book
		if err = rows.StructScan(&book); err != nil {
			return nil, fmt.Errorf("StructcScan: %w", err)
		}
		books = append(books, book)
	}
	return books, nil
}
func (st *RepoStorage) UpdateRepo(ctx context.Context, repo entity.Repo) error {
	var r entity.Repo
	if err := st.db.QueryRowxContext(ctx, updateRepo,
		&repo.Name,
		&repo.Visibility,
		&repo.Desc,
		&repo.ID,
	).StructScan(&r); err != nil {
		return fmt.Errorf("QueryRowxContext: %w", err)
	}
	return nil
}
func (st *RepoStorage) DeleteRepo(ctx context.Context, repoID int64) error {
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
