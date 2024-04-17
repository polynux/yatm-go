package main

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
	"polynux/yatm/handlers"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	templates, err := parseTemplates("views")
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Renderer = &TemplateRenderer{
		templates: templates,
	}

	e.Static("/static", "static")

	e.GET("/", handlers.HomeHandler)
	e.GET("/search", handlers.SearchHandler)

	e.Logger.Fatal(e.Start(":3000"))
}

func populateTemplates(dir string, patterns []string) *template.Template {
	templates := template.New("")
	templates = templates.Funcs(template.FuncMap{
		"assets": assets,
	})

	for _, pattern := range patterns {
		templates = template.Must(templates.ParseGlob(filepath.Join(dir, pattern)))
	}

	return templates
}

func parseTemplates(root string) (*template.Template, error) {
	templates := template.New("")
	templates = templates.Funcs(template.FuncMap{
		"assets": assets,
		"svg":    loadSvg,
	})

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".html") {
			_, err = templates.ParseFiles(path)
			if err != nil {
				return err
			}
		}

		return nil
	}

	err := filepath.Walk(root, walkFn)
	if err != nil {
		return nil, err
	}

	return templates, nil
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func assets(path string) string {
	return "/static/" + path
}

func loadSvg(name, class string) (template.HTML, error) {
	path := filepath.Join("static", "svg", name+".svg")
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	content := string(data)
	if class != "" {
		content = strings.ReplaceAll(content, "<svg ", "<svg class=\""+class+"\" ")
	}

	return template.HTML(content), nil
}
