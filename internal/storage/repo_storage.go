package storage

import (
	"context"
	"fmt"
	"github.com/VeneLooool/BookHub/internal/entity"
	"strconv"
)

type RepoStorageAbs struct {
	db    RepoStorage
	cache Cache[string, entity.Repo]
}

func NewRepoStorageAbs(db RepoStorage, ch Cache[string, entity.Repo]) *RepoStorageAbs {
	return &RepoStorageAbs{
		db:    db,
		cache: ch,
	}
}

func (rsa *RepoStorageAbs) CreateRepo(ctx context.Context, userID int64, repo entity.Repo) (ID int64, err error) {
	ID, err = rsa.db.CreateRepo(ctx, userID, repo)
	if err != nil {
		return 0, fmt.Errorf("db CreateRepo: %w", err)
	}

	key := rsa.idToCacheKey(ID)
	if err = rsa.cache.Set(key, repo); err != nil {
		//TODO add logger
	}
	return ID, nil
}
func (rsa *RepoStorageAbs) GetReposForUser(ctx context.Context, userID int64) ([]entity.Repo, error) {
	return rsa.db.GetReposForUser(ctx, userID)
}
func (rsa *RepoStorageAbs) GetRepo(ctx context.Context, ID int64) (repo entity.Repo, err error) {
	key := rsa.idToCacheKey(ID)
	exist, err := rsa.cache.Exist(key)
	if err != nil {
		//TODO add logger
	}
	if exist {
		repo, err = rsa.cache.Get(key)
		if err == nil {
			return repo, nil
		}
		//TODO add logger
	}
	repo, err = rsa.db.GetRepo(ctx, ID)
	if err != nil {
		return entity.Repo{}, fmt.Errorf("db GetRepo: %w", err)
	}
	return repo, nil
}
func (rsa *RepoStorageAbs) UpdateRepo(ctx context.Context, repo entity.Repo) error {
	err := rsa.db.UpdateRepo(ctx, repo)
	if err != nil {
		return fmt.Errorf("db UpdateRepo: %w", err)
	}

	key := rsa.idToCacheKey(repo.ID)
	err = rsa.cache.Update(key, repo)
	if err != nil {
		//TODO add logger
	}
	return nil
}
func (rsa *RepoStorageAbs) DeleteRepo(ctx context.Context, ID int64) error {
	err := rsa.db.DeleteRepo(ctx, ID)
	if err != nil {
		return fmt.Errorf("db DeleteRepo: %w", err)
	}

	key := rsa.idToCacheKey(ID)
	err = rsa.cache.Delete(key)
	if err != nil {
		//TODO add logger
	}
	return nil
}
func (rsa *RepoStorageAbs) DeleteBookFromRepo(ctx context.Context, RepoID, bookID int64) error {
	return rsa.db.DeleteBookFromRepo(ctx, RepoID, bookID)
}
func (rsa *RepoStorageAbs) idToCacheKey(ID int64) string {
	return strconv.Itoa(int(ID))
}
