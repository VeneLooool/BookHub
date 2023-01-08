package grpc

import (
	"github.com/VeneLooool/BookHub/internal/entity"
	desc "github.com/VeneLooool/BookHub/internal/pb"
)

func transformUserFromGrpcToEntity(user *desc.User) entity.User {
	return entity.User{
		ID:       user.GetId(),
		Name:     user.GetName(),
		UserName: user.GetUsername(),
		Password: user.GetPassword(),
		Desc:     user.GetDescription(),
	}
}

func transformUserFromEntityToGrpc(user entity.User) *desc.User {
	return &desc.User{
		Id:          user.ID,
		Name:        user.Name,
		Username:    user.UserName,
		Password:    user.Password,
		Description: user.Desc,
	}
}

func transformRepoFromGrpcToEntity(repo *desc.Repo) entity.Repo {
	return entity.Repo{
		ID:         repo.GetId(),
		Name:       repo.GetName(),
		Visibility: repo.GetVisibility(),
		Desc:       repo.GetDescription(),
		UserID:     repo.GetUserId(),
	}
}

func transformRepoFromEntityToGrpc(repo entity.Repo) *desc.Repo {
	return &desc.Repo{
		Id:          repo.ID,
		Name:        repo.Name,
		Visibility:  repo.Visibility,
		Description: repo.Desc,
		UserId:      repo.UserID,
	}
}

func transformBookFromGrpcToEntity(book *desc.Book) entity.Book {
	return entity.Book{
		ID:          book.GetId(),
		Title:       book.GetTitle(),
		Author:      book.GetAuthor(),
		NumberPages: book.GetCurrentPage(),
		Desc:        book.GetDescription(),
		File:        entity.File{Type: entity.PDF},
		Image:       entity.File{Type: entity.Image},
	}
}

func transformBookFromEntityToGrpc(book entity.Book) *desc.Book {
	return &desc.Book{
		Id:          book.ID,
		Title:       book.Title,
		Author:      book.Author,
		NumberPages: book.NumberPages,
		CurrentPage: book.CurrentPage,
		Description: book.Desc,
	}
}
