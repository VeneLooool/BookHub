package usecase

import (
	"bookhub/internal/entity"
	"bookhub/internal/storage"
	"context"
	"fmt"
)

type RepoService struct {
	storage storage.RepoStorage
}

func NewRepoService(repoStorage storage.RepoStorage) *RepoService {
	return &RepoService{
		storage: repoStorage,
	}
}

func (rs *RepoService) CreateRepo(ctx context.Context, userID int64, repo entity.Repo) (ID int64, err error) {
	ID, err = rs.storage.CreateRepo(ctx, userID, repo)
	if err != nil {
		return 0, fmt.Errorf("CreateRepo: %w", err)
	}
	return ID, nil
}
func (rs *RepoService) GetRepo(ctx context.Context, ID int64) (repo entity.Repo, err error) {
	repo, err = rs.storage.GetRepo(ctx, ID)
	if err != nil {
		return entity.Repo{}, fmt.Errorf("GetRepo: %w", err)
	}
	return repo, nil
}
func (rs *RepoService) GetBooksForRepo(ctx context.Context, repoID int64) (books []entity.Book, err error) {
	books, err = rs.storage.GetBooksForRepo(ctx, repoID)
	if err != nil {
		return nil, fmt.Errorf("GetBookForRepo: %w", err)
	}
	return books, err
}
func (rs *RepoService) updateRepo(oldRepo, newRepo entity.Repo) entity.Repo {
	if newRepo.Visibility != "" {
		oldRepo.Visibility = newRepo.Visibility
	}
	if newRepo.Name != "" {
		oldRepo.Name = newRepo.Name
	}
	if newRepo.Desc != "" {
		oldRepo.Desc = newRepo.Desc
	}
	return oldRepo
}

func (rs *RepoService) UpdateRepo(ctx context.Context, newRepo entity.Repo) (repo entity.Repo, err error) {
	repo, err = rs.GetRepo(ctx, newRepo.ID)
	if err != nil {
		return entity.Repo{}, fmt.Errorf("GetRepo: %w", err)
	}

	repo = rs.updateRepo(repo, newRepo)
	err = rs.storage.UpdateRepo(ctx, repo)
	if err != nil {
		return entity.Repo{}, fmt.Errorf("UpdateRepo: %w", err)
	}
	return repo, nil
}
func (rs *RepoService) DeleteRepo(ctx context.Context, ID int64) (err error) {
	if err = rs.storage.DeleteRepo(ctx, ID); err != nil {
		return fmt.Errorf("DeleteRepo: %w", err)
	}
	return nil
}

func (rs *RepoService) DeleteBookFromRepo(ctx context.Context, RepoID, bookID int64) (err error) {
	if err = rs.storage.DeleteBookFromRepo(ctx, RepoID, bookID); err != nil {
		return fmt.Errorf("DeleteRepo: %w", err)
	}
	return nil
}
