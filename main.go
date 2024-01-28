package main

import (
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
		followersHTML, followingsHTML, err := core.AnalyzeInstagramFollowersHTML(c)
		if err != nil {
			return c.Render(http.StatusBadRequest, "main.html", map[string]interface{}{
				"errorMessage": "Oops! Something went wrong. Please check your file uploads and try again. Refer to the instructions above.",
			})
		}

		var results core.PageVariables

		if len(followersHTML) > 0 && len(followingsHTML) > 0 {
			usersImNotFollowingBackHTML := core.CheckFollowStatusHTML(followingsHTML, followersHTML)
			usersNotFollowingMeBackHTML := core.CheckFollowStatusHTML(followersHTML, followingsHTML)

			results.UsersNotFollowingMeBack = usersNotFollowingMeBackHTML
			results.UsersImNotFollowingBack = usersImNotFollowingBackHTML
		} else {
			usersImNotFollowingBackJSON, usersNotFollowingMeBackJSON, err := core.AnalyzeInstagramFollowersJSON(c)
			if err != nil {
				return c.Render(http.StatusBadRequest, "main.html", map[string]interface{}{
					"ErrorMessage": "Oops! Something went wrong. Please check your file uploads and try again. Refer to the instructions above.",
				})
			}

			results.UsersNotFollowingMeBack = usersNotFollowingMeBackJSON
			results.UsersImNotFollowingBack = usersImNotFollowingBackJSON
		}

		return c.Render(http.StatusOK, "main.html", results)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
