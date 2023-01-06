package grpc

import (
	"context"
	desc "github.com/VeneLooool/BookHub/internal/pb"
	"github.com/VeneLooool/BookHub/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RepoService struct {
	uc service.RepoUseCase
}

func (rs *RepoService) CreateRepo(ctx context.Context, in *desc.CreateRepoReq, opts ...grpc.CallOption) (*desc.Repo, error) {
	return nil, nil
}
func (rs *RepoService) GetReposForUser(ctx context.Context, in *desc.GetReposForUserReq, opts ...grpc.CallOption) (*desc.GetReposForUserResp, error) {
	return nil, nil
}
func (rs *RepoService) UpdateRepo(ctx context.Context, in *desc.UpdateRepoReq, opts ...grpc.CallOption) (*desc.Repo, error) {
	return nil, nil
}
func (rs *RepoService) GetRepo(ctx context.Context, in *desc.GetRepoReq, opts ...grpc.CallOption) (*desc.Repo, error) {
	return nil, nil
}
func (rs *RepoService) DeleteRepo(ctx context.Context, in *desc.DeleteRepoReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}
func (rs *RepoService) DeleteBookFromRepo(ctx context.Context, in *desc.DeleteBookFromRepoReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}
