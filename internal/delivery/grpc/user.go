package grpc

import (
	"context"
	"github.com/VeneLooool/BookHub/internal/entity"
	desc "github.com/VeneLooool/BookHub/internal/pb"
	"github.com/VeneLooool/BookHub/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	uc service.UserUseCase
}

func (us *UserService) CreateUser(ctx context.Context, in *desc.CreateUserReq, opts ...grpc.CallOption) (*desc.User, error) {
	userId, err := us.uc.CreateUser(ctx, transformUserFromGrpcToEntity(in.GetUser()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return transformUserFromEntityToGrpc(entity.User{ID: userId}), nil
}
func (us *UserService) GetUser(ctx context.Context, in *desc.GetUserReq, opts ...grpc.CallOption) (*desc.User, error) {
	user, err := us.uc.GetUser(ctx, in.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return transformUserFromEntityToGrpc(user), nil
}
func (us *UserService) UpdateUser(ctx context.Context, in *desc.UpdateUserReq, opts ...grpc.CallOption) (*desc.User, error) {
	newUser, err := us.uc.UpdateUser(ctx, transformUserFromGrpcToEntity(in.GetUser()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return transformUserFromEntityToGrpc(newUser), nil
}
func (us *UserService) DeleteUser(ctx context.Context, in *desc.DeleteUserReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	err := us.uc.DeleteUser(ctx, in.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return nil, nil
}
