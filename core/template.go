package core

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
)

var mainTmpl = `{{ define "main" }}{{ template "base" . }}{{ end }}`

type templateRenderer struct {
	templates map[string]*template.Template
}

func NewTemplateRenderer(layoutsDir, viewsDir string) *templateRenderer {
	r := &templateRenderer{}
	r.templates = make(map[string]*template.Template)
	r.Load(layoutsDir, viewsDir)
	return r
}

func (t *templateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		c.Logger().Fatalf("the template %s does not exist", name)
		return fmt.Errorf("the template %s does not exist", name)
	}

	return tmpl.ExecuteTemplate(w, "base", data)
}

func (t *templateRenderer) Load(layoutsDir, templatesDir string) {
    layouts, err := filepath.Glob(layoutsDir + "/*.html")
    if err != nil {
        log.Fatal(err)
    }

    includes, err := filepath.Glob(templatesDir + "/*.html")
    if err != nil {
        log.Fatal(err)
    }

    mainTemplate := template.New("main")

    mainTemplate, err = mainTemplate.Parse(mainTmpl)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range includes {
        fileName := filepath.Base(file)
        files := append(layouts, file)
        t.templates[fileName], err = mainTemplate.Clone()

        if err != nil {
            log.Fatal(err)
        }

        t.templates[fileName] = template.Must(t.templates[fileName].ParseFiles(files...))
    }
}