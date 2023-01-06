package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/VeneLooool/BookHub/internal/entity"
	"github.com/VeneLooool/BookHub/internal/storage"
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
	result, err := st.db.ExecContext(ctx, createUser, &user.Name, &user.UserName, &user.Password, &user.Desc)
	if err != nil {
		return 0, fmt.Errorf("QueryRowxContext: %w", err)
	}

	ID, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("LastInsertId: %w", err)
	}
	return u.ID, nil
}
func (st *UserStorage) GetUser(ctx context.Context, ID int64) (user entity.User, err error) {
	if err = st.db.GetContext(ctx, &user, getUser, &ID); err != nil {
		return entity.User{}, fmt.Errorf("GetContext: %w", err)
	}
	return user, nil
}

func (st *UserStorage) UpdateUser(ctx context.Context, user entity.User) (err error) {
	result, err := st.db.ExecContext(ctx, updateUser, &user.Name, &user.UserName,
		&user.Password, &user.Desc, &user.ID)
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
