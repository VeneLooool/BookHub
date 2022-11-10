package http

import (
	"github.com/labstack/echo/v4"
)

type UserHandlers interface {
	CreateUser() echo.HandlerFunc
	GetUser() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
}

type RepoHandlers interface {
	CreateRepo() echo.HandlerFunc
	GetReposForUser() echo.HandlerFunc
	GetRepo() echo.HandlerFunc
	UpdateRepo() echo.HandlerFunc
	DeleteRepo() echo.HandlerFunc
	DeleteBookFromRepo() echo.HandlerFunc
}

type BookHandlers interface {
	CreateBook() echo.HandlerFunc
	GetBook() echo.HandlerFunc
	GetBookFile() echo.HandlerFunc
	GetBookImage() echo.HandlerFunc
	GetBooksForRepo() echo.HandlerFunc
	UpdateBook() echo.HandlerFunc
	DeleteBook() echo.HandlerFunc
}
