package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func HandleHome(c echo.Context) error {
    return c.Redirect(http.StatusMovedPermanently, "https://github.com/samperlmutter")
}

func HandleResume(c echo.Context) error {
    return c.File("assets/resume.pdf")
}
