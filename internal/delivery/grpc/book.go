package grpc

import (
	"context"
	desc "github.com/VeneLooool/BookHub/internal/pb"
	"github.com/VeneLooool/BookHub/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BookService struct {
	service.BookUseCase
}

func (bs *BookService) CreateBook(ctx context.Context, in *desc.CreateBookReq, opts ...grpc.CallOption) (*desc.Book, error) {
	return nil, nil
}
func (bs *BookService) GetBooksForRepo(ctx context.Context, in *desc.GetBooksForRepoReq, opts ...grpc.CallOption) (*desc.GetBooksForRepoResp, error) {
	return nil, nil
}
func (bs *BookService) UpdateBook(ctx context.Context, in *desc.UpdateBookReq, opts ...grpc.CallOption) (*desc.Book, error) {
	return nil, nil
}
func (bs *BookService) GetBook(ctx context.Context, in *desc.GetBookReq, opts ...grpc.CallOption) (*desc.Book, error) {
	return nil, nil
}
func (bs *BookService) GetBookImage(ctx context.Context, in *desc.GetBookImageReq, opts ...grpc.CallOption) (*desc.File, error) {
	return nil, nil
}
func (bs *BookService) GetBookFile(ctx context.Context, in *desc.GetBookFileReq, opts ...grpc.CallOption) (*desc.File, error) {
	return nil, nil
}
func (bs *BookService) DeleteBook(ctx context.Context, in *desc.DeleteBookReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}
