package handlers

import (
	"bookhub/internal/config"
	intHttp "bookhub/internal/delivery/http"
	"bookhub/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type repoHandlers struct {
	cfg    *config.Config
	repoUC service.RepoUseCase
	//logger
}

func NewRepoHandlers(cfg *config.Config, repoUC service.RepoUseCase) intHttp.RepoHandlers {
	return &repoHandlers{
		repoUC: repoUC,
		cfg:    cfg,
	}
}
func (h *repoHandlers) CreateRepo() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusCreated, nil)
	}
}
func (h *repoHandlers) GetReposForUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
func (h *repoHandlers) GetRepo() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
func (h *repoHandlers) UpdateRepo() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
func (h *repoHandlers) DeleteRepo() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
func (h *repoHandlers) DeleteBookFromRepo() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
