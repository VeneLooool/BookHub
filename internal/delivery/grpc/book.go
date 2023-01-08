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

type BookService struct {
	uc service.BookUseCase
}

func (bs *BookService) CreateBook(ctx context.Context, in *desc.CreateBookReq, opts ...grpc.CallOption) (*desc.Book, error) {
	book := transformBookFromGrpcToEntity(in.GetBook())
	book.File.File = in.GetFile().GetFile()
	book.Image.File = in.GetImage().GetFile()

	bookId, err := bs.uc.CreateBook(ctx, in.GetRepoId(), book)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return transformBookFromEntityToGrpc(entity.Book{ID: bookId}), nil
}
func (bs *BookService) GetBooksForRepo(ctx context.Context, in *desc.GetBooksForRepoReq, opts ...grpc.CallOption) (*desc.GetBooksForRepoResp, error) {
	books, err := bs.uc.GetBooksForRepo(ctx, in.GetRepoId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	grpcBooks := make([]*desc.Book, 0, len(books))
	for _, book := range books {
		grpcBooks = append(grpcBooks, transformBookFromEntityToGrpc(book))
	}
	return &desc.GetBooksForRepoResp{Books: grpcBooks}, nil
}
func (bs *BookService) UpdateBook(ctx context.Context, in *desc.UpdateBookReq, opts ...grpc.CallOption) (*desc.Book, error) {
	book, err := bs.uc.UpdateBook(ctx, transformBookFromGrpcToEntity(in.GetBook()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return transformBookFromEntityToGrpc(book), nil
}
func (bs *BookService) GetBook(ctx context.Context, in *desc.GetBookReq, opts ...grpc.CallOption) (*desc.Book, error) {
	book, err := bs.uc.GetBook(ctx, in.GetBookId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return transformBookFromEntityToGrpc(book), nil
}
func (bs *BookService) GetBookImage(ctx context.Context, in *desc.GetBookImageReq, opts ...grpc.CallOption) (*desc.File, error) {
	bookImage, err := bs.uc.GetBookImage(ctx, in.GetBookId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &desc.File{File: bookImage.File, FileType: string(bookImage.Type)}, nil
}
func (bs *BookService) GetBookFile(ctx context.Context, in *desc.GetBookFileReq, opts ...grpc.CallOption) (*desc.File, error) {
	bookFile, err := bs.uc.GetBookFile(ctx, in.GetBookId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &desc.File{File: bookFile.File, FileType: string(bookFile.Type)}, nil
}
func (bs *BookService) DeleteBook(ctx context.Context, in *desc.DeleteBookReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	err := bs.uc.DeleteBook(ctx, in.GetBookId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return nil, nil
}
