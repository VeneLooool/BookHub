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
func (st *UserStorage) GetReposForUser(ctx context.Context, userID int64) ([]entity.Repo, error) {
	return nil, nil
}
func (st *UserStorage) UpdateUser(context.Context, entity.User) error {
	return nil
}
func (st *UserStorage) DeleteUser(context.Context, int64) error {
	return nil
}
