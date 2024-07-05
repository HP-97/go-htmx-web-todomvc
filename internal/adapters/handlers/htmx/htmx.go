package htmx

import (
	"embed"
	"html/template"

	"github.com/HP-97/go-htmx-web-todomvc/internal/core/ports"
)

//go:embed templates
var tmplFs embed.FS

type HTMXHandler struct {
	srv ports.TodoService
	tmpl *template.Template
}

func NewHTMXHandler(srv ports.TodoService) (*HTMXHandler, error) {
	// Parse templates
	funcs := template.FuncMap(template.FuncMap{
		"attr": func(s string) template.HTMLAttr {
			return template.HTMLAttr(s)
		},
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	})

	tmpl, err := template.New("todo").Funcs(funcs).ParseFS(tmplFs, "templates/*.html")
	if err != nil {
		return nil, err
	}

	return &HTMXHandler{
		srv: srv,
		tmpl: tmpl,
	}, nil
}
