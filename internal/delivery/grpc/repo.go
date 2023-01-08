package grpc

import (
	"context"
	"github.com/VeneLooool/BookHub/internal/entity"
	desc "github.com/VeneLooool/BookHub/internal/pb"
	"github.com/VeneLooool/BookHub/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RepoService struct {
	uc service.RepoUseCase
}

func (rs *RepoService) CreateRepo(ctx context.Context, in *desc.CreateRepoReq) (*desc.Repo, error) {
	repoId, err := rs.uc.CreateRepo(ctx, in.GetUserId(), transformRepoFromGrpcToEntity(in.GetRepo()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return transformRepoFromEntityToGrpc(entity.Repo{ID: repoId}), nil
}
func (rs *RepoService) GetReposForUser(ctx context.Context, in *desc.GetReposForUserReq) (*desc.GetReposForUserResp, error) {
	repos, err := rs.uc.GetReposForUser(ctx, in.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	grpcRepos := make([]*desc.Repo, 0, len(repos))
	for _, repo := range repos {
		grpcRepos = append(grpcRepos, transformRepoFromEntityToGrpc(repo))
	}
	return &desc.GetReposForUserResp{Repos: grpcRepos}, nil
}
func (rs *RepoService) UpdateRepo(ctx context.Context, in *desc.UpdateRepoReq) (*desc.Repo, error) {
	repo, err := rs.uc.UpdateRepo(ctx, transformRepoFromGrpcToEntity(in.GetRepo()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return transformRepoFromEntityToGrpc(repo), nil
}
func (rs *RepoService) GetRepo(ctx context.Context, in *desc.GetRepoReq) (*desc.Repo, error) {
	repo, err := rs.uc.GetRepo(ctx, in.GetRepoId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return transformRepoFromEntityToGrpc(repo), nil
}
func (rs *RepoService) DeleteRepo(ctx context.Context, in *desc.DeleteRepoReq) (*emptypb.Empty, error) {
	err := rs.uc.DeleteRepo(ctx, in.GetRepoId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (rs *RepoService) DeleteBookFromRepo(ctx context.Context, in *desc.DeleteBookFromRepoReq) (*emptypb.Empty, error) {
	err := rs.uc.DeleteBookFromRepo(ctx, in.GetRepoId(), in.GetBookId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
