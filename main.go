package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

type PageVariables struct {
	Message string
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
	
	e.Renderer = NewTemplateRenderer("public/views/*.html")

	e.GET("/", handler)
	e.GET("/*", notFound)

	e.Logger.Fatal(e.Start(":8080"))
}

func NewTemplateRenderer(dir string) *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseGlob(dir)),
	}
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}