package main

import (
	"context"
	"fmt"

	"github.com/HP-97/go-htmx-web-todomvc/http"
)

// Use this reddit post to understand Domain-Driven Design
// This is a GREAT example on how to set up a Golang repo https://github.com/benbjohnson/wtf
// https://www.reddit.com/r/golang/comments/15rd0xu/looking_for_go_projects_that_applies_hexagonal/
// DEPRECATED: Use this git repo as a baseline https://github.com/LordMoMA/Hexagonal-Architecture
func main() {

	// TODO: Initialise a config file
	// TODO: Initialise a repository struct
	// TODO: Initialise a service struct
	// TODO: Initialise a handler struct
	fmt.Println("hello world!")
}

type Main struct {
	HTTPServer *http.Server
}

func NewMain() *Main {
	return &Main {
		HTTPServer: http.NewServer(),
	}
}

func (m *Main) Run(ctx context.Context) (err error) {

}

