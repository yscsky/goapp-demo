package main

import (
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &home{})
	app.RunWhenOnBrowser()

	h := &app.Handler{
		Name: "Bulma Demo",
		Styles: []string{
			"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css",
		},
	}
	http.Handle("/", h)
	http.ListenAndServe(":8080", nil)
}
