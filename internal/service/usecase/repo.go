package usecase

import (
	"bookhub/internal/entity"
	"context"
)

type RepoStorage interface {
	CreateRepo(context.Context, entity.Repo) (int64, error)
	GetRepo(context.Context, int64) (entity.Repo, error)
	UpdateRepo(context.Context, entity.Repo) error
	DeleteRepo(context.Context, int64) error
}

type RepoService struct {
	storage RepoStorage
}

func NewRepoService(repoStorage RepoStorage) *RepoService {
	return &RepoService{
		storage: repoStorage,
	}
}

func (rs *RepoService) CreateRepo(ctx context.Context, repo entity.Repo) (int64, error) {
	return rs.storage.CreateRepo(ctx, repo)
}
func (rs *RepoService) GetRepo(ctx context.Context, ID int64) (entity.Repo, error) {
	return rs.storage.GetRepo(ctx, ID)
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

func (rs *RepoService) UpdateRepo(ctx context.Context, newRepo entity.Repo) (entity.Repo, error) {
	repo, err := rs.GetRepo(ctx, newRepo.ID)
	if err != nil {
		return entity.Repo{}, err
	}

	repo = rs.updateRepo(repo, newRepo)
	err = rs.storage.UpdateRepo(ctx, repo)
	if err != nil {
		return entity.Repo{}, err
	}
	return repo, nil
}
func (rs *RepoService) DeleteRepo(ctx context.Context, ID int64) error {
	return rs.storage.DeleteRepo(ctx, ID)
}
