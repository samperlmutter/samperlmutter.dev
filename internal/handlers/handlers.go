package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct { }

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandleHome(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "https://github.com/samperlmutter")
}

func (h *Handler) HandleResume(c echo.Context) error {
	return c.File("assets/resume.pdf")
}
