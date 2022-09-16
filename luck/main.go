package main

import (
	"net/http"

	"goapp-demos/luck/route"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &route.Home{})
	app.Route("/select", &route.GameSelect{})
	app.RunWhenOnBrowser()

	h := &app.Handler{
		Name: "Luck",
		Icon: app.Icon{
			Default: "/web/icon.png",
		},
		Styles: []string{
			"/web/luck.css",
		},
		CacheableResources: []string{
			"/web/bg.jpg",
		},
	}
	http.Handle("/", h)
	http.ListenAndServe(":8080", nil)
}
