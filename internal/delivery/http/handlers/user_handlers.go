package handlers

import (
	"github.com/VeneLooool/BookHub/internal/config"
	intHttp "github.com/VeneLooool/BookHub/internal/delivery/http"
	"github.com/VeneLooool/BookHub/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type userHandlers struct {
	cfg    *config.Config
	userUC service.UserUseCase
	//logger
}

func NewUserHandlers(cfg *config.Config, userUC service.UserUseCase) intHttp.UserHandlers {
	return &userHandlers{
		userUC: userUC,
		cfg:    cfg,
	}
}

func (h *userHandlers) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusCreated, nil)
	}
}
func (h *userHandlers) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
func (h *userHandlers) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
func (h *userHandlers) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
