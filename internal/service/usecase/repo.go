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
