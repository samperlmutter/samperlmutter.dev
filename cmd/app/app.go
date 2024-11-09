package app

import (
    "samperlmutter.dev/internal/handlers"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func Serve() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", handlers.HandleHome)
    e.GET("/resume", handlers.HandleResume)

    e.Logger.Fatal(e.Start(":80"))
} 