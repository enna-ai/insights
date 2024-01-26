package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
	"github.com/enna-ai/instacheck/core"
)

type PageVariables struct {
	Message               string
	UsersNotFollowingBack []string
	NotFollowingUsers     []string
}

func handler(c echo.Context) error {
	pageVariables := PageVariables{
		Message: "GO",
	}

	return c.Render(http.StatusOK, "main.html", pageVariables)
}

func notFound(c echo.Context) error {
	return c.Render(http.StatusNotFound, "404.html", nil)
}

func main() {
	e := echo.New()

	e.Renderer = core.NewTemplateRenderer("templates/layouts", "templates/views")

	e.GET("/", handler)
	e.GET("/*", notFound)

	e.Logger.Fatal(e.Start(":8080"))
}
