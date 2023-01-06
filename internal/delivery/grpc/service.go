package grpc

import (
	desc "github.com/VeneLooool/BookHub/internal/pb"
)

type Service struct {
	UserService
	RepoService
	BookService
	desc.UnimplementedBookHubServiceServer
}

func NewService() *Service {
	return &Service{}
}
