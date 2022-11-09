package storage

import (
	"bookhub/internal/entity"
	"context"
	"fmt"
	"strconv"
)

type UserStorageAbs struct {
	db    UserStorage
	cache Cache[string, entity.User]
}

func NewUserStorageAbs(db UserStorage, ch Cache[string, entity.User]) *UserStorageAbs {
	return &UserStorageAbs{
		db:    db,
		cache: ch,
	}
}

func (usa *UserStorageAbs) CreateUser(ctx context.Context, user entity.User) (ID int64, err error) {
	ID, err = usa.db.CreateUser(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("db CreateUser: %w", err)
	}

	key := usa.idToCacheKey(ID)
	if err = usa.cache.Set(key, user); err != nil {
		//TODO add logger
	}
	return ID, nil
}
func (usa *UserStorageAbs) GetUser(ctx context.Context, ID int64) (user entity.User, err error) {
	key := usa.idToCacheKey(ID)
	exist, err := usa.cache.Exist(key)
	if err != nil {
		//TODO add logger
	}
	if exist {
		user, err = usa.cache.Get(key)
		if err == nil {
			return user, nil
		}
		//TODO add logger
	}

	user, err = usa.db.GetUser(ctx, ID)
	if err != nil {
		return entity.User{}, fmt.Errorf("db GetUser: %w", err)
	}
	return user, err
}
func (usa *UserStorageAbs) UpdateUser(ctx context.Context, user entity.User) error {
	err := usa.db.UpdateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("db UpdateUser: %w", err)
	}

	key := usa.idToCacheKey(user.ID)
	if err = usa.cache.Update(key, user); err != nil {
		//TODO add logger
	}
	return nil
}
func (usa *UserStorageAbs) DeleteUser(ctx context.Context, ID int64) error {
	err := usa.db.DeleteUser(ctx, ID)
	if err != nil {
		return err
	}

	key := usa.idToCacheKey(ID)
	if err = usa.cache.Delete(key); err != nil {
		//TODO add logger
	}
	return nil
}

func (usa *UserStorageAbs) idToCacheKey(ID int64) string {
	return strconv.Itoa(int(ID))
}
