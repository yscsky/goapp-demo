package main

import (
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &hello{})
	app.Route("/city", &city{})
	app.RunWhenOnBrowser()

	h := &app.Handler{
		Name: "App Demo",
		Styles: []string{
			"/web/hello.css",
			"/web/city.css",
		},
		CacheableResources: []string{
			"/web/space.jpg",
			"/web/beijing.jpg",
			"/web/paris.jpg",
			"/web/sf.jpg",
		},
	}
	http.Handle("/", h)
	http.ListenAndServe(":8080", nil)
}
