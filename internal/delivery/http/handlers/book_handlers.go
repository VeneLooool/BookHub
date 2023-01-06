package handlers

import (
	"github.com/VeneLooool/BookHub/internal/config"
	intHttp "github.com/VeneLooool/BookHub/internal/delivery/http"
	"github.com/VeneLooool/BookHub/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type bookHandlers struct {
	cfg    *config.Config
	bookUC service.BookUseCase
	//logger
}

func NewBookHandlers(cfg *config.Config, bookUC service.BookUseCase) intHttp.BookHandlers {
	return &bookHandlers{
		bookUC: bookUC,
		cfg:    cfg,
	}
}
func (h *bookHandlers) CreateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusCreated, nil)
	}
}
func (h *bookHandlers) GetBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
func (h *bookHandlers) GetBookFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
func (h *bookHandlers) GetBookImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
func (h *bookHandlers) GetBooksForRepo() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
func (h *bookHandlers) UpdateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
func (h *bookHandlers) DeleteBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
