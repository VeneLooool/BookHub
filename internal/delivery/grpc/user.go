package grpc

import (
	"context"
	desc "github.com/VeneLooool/BookHub/internal/pb"
	"github.com/VeneLooool/BookHub/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	service.UserUseCase
}

func (us *UserService) CreateUser(ctx context.Context, in *desc.CreateUserReq, opts ...grpc.CallOption) (*desc.User, error) {
	return nil, nil
}
func (us *UserService) GetUser(ctx context.Context, in *desc.GetUserReq, opts ...grpc.CallOption) (*desc.User, error) {
	return nil, nil
}
func (us *UserService) UpdateUser(ctx context.Context, in *desc.UpdateUserReq, opts ...grpc.CallOption) (*desc.User, error) {
	return nil, nil
}
func (us *UserService) DeleteUser(ctx context.Context, in *desc.DeleteUserReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}
