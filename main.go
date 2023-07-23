package main

import (
	"os"

	"github.com/amanthakur0001/github-actions/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", api.Hello)

	e.GET("/health", api.Health)

	e.GET("/print", api.Print)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
