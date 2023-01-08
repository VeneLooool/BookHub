package grpc

import (
	desc "github.com/VeneLooool/BookHub/internal/pb"
	"github.com/VeneLooool/BookHub/internal/service"
)

type Service struct {
	UserService
	RepoService
	BookService
	desc.UnimplementedBookHubServiceServer
}

func NewService(userUseCase service.UserUseCase, repoUseCase service.RepoUseCase, bookUseCase service.BookUseCase) *Service {
	return &Service{
		UserService: UserService{
			uc: userUseCase,
		},
		RepoService: RepoService{
			uc: repoUseCase,
		},
		BookService: BookService{
			uc: bookUseCase,
		},
	}
}
