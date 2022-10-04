package postgres

import (
	"bookhub/internal/entity"
	"bookhub/internal/storage"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserStorage struct {
	db *sqlx.DB
}

func NewUserStorage(db *sqlx.DB) storage.UserStorage {
	return &UserStorage{db: db}
}

func (st *UserStorage) CreateUser(ctx context.Context, user entity.User) (ID int64, err error) {
	var u entity.User

	if err = st.db.QueryRowxContext(
		ctx,
		createUser,
		&user.ID,
		&user.Name,
		&user.UserName,
		&user.Password,
		&user.Desc,
	).Scan(&u); err != nil {
		return 0, fmt.Errorf("QueryRowxContext: %w", err)
	}
	return u.ID, err
}
func (st *UserStorage) GetUser(ctx context.Context, ID int64) (user entity.User, err error) {
	if err = st.db.GetContext(ctx, &user, getUser, &ID); err != nil {
		return entity.User{}, fmt.Errorf("GetContext: %w", err)
	}
	return user, nil
}
func (st *UserStorage) GetReposForUser(ctx context.Context, userID int64) (repos []entity.Repo, err error) {
	rows, err := st.db.QueryxContext(ctx, getReposForUser, &userID)
	if err != nil {
		return nil, fmt.Errorf("QueryxContext: %w", err)
	}
	for rows.Next() {
		var repo entity.Repo
		if err = rows.StructScan(&repo); err != nil {
			return nil, fmt.Errorf("StructScan: %w", err)
		}
		repos = append(repos, repo)
	}
	return repos, nil
}
func (st *UserStorage) UpdateUser(ctx context.Context, user entity.User) (err error) {
	var u entity.User
	if err = st.db.QueryRowxContext(ctx, updateUser,
		&user.Name,
		&user.UserName,
		&user.Password,
		&user.Desc,
		&user.ID,
	).StructScan(&u); err != nil {
		return fmt.Errorf("QueryRowxContext: %w", err)
	}
	return nil
}
func (st *UserStorage) DeleteUser(ctx context.Context, ID int64) (err error) {
	result, err := st.db.ExecContext(ctx, deleteUser, ID)
	if err != nil {
		return fmt.Errorf("ExecContext: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("RowsAffected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("RowsAffceted: %w", entity.ErrUserNotFound)
	}
	return nil
}
