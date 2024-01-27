package main

import (
	"fmt"
	"net/http"

	"github.com/enna-ai/instacheck/core"
	"github.com/labstack/echo/v4"
)

func handler(c echo.Context) error {
	return c.Render(http.StatusOK, "main.html", nil)
}

func notFound(c echo.Context) error {
	return c.Render(http.StatusNotFound, "404.html", nil)
}

func main() {
	e := echo.New()

	e.Renderer = core.NewTemplateRenderer("templates/layouts", "templates/views")

	e.GET("/", handler)
	e.GET("/*", notFound)
	e.POST("/upload", func(c echo.Context) error {
		usersImNotFollowingBack, usersNotFollowingMeBack, err := core.ParseJSONFiles(c)
		if err != nil {
			fmt.Printf("Error %s\n", err)
			return err
		}

		results := core.PageVariables{
			UsersNotFollowingMeBack: usersNotFollowingMeBack,
			UsersImNotFollowingBack: usersImNotFollowingBack,
		}

		return c.Render(http.StatusOK, "main.html", results)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
