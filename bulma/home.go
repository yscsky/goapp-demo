package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type home struct {
	app.Compo
}

func (h *home) Render() app.UI {
	return app.Div().Style("padding", "20px").Body(
		app.Button().Class("button").Body(app.Text("button")),
		app.Button().Class("button is-primary").Body(app.Text("button")),
		app.Button().Class("button is-link").Body(app.Text("button")),
		app.Button().Class("button is-info").Body(app.Text("button")),
		app.Button().Class("button is-success").Body(app.Text("button")),
		app.Button().Class("button is-warning").Body(app.Text("button")),
		app.Button().Class("button is-danger").Body(app.Text("button")),
		app.Button().Class("button is-small").Body(app.Text("button")),
		app.Button().Class("button is-medium").Body(app.Text("button")),
		app.Button().Class("button is-large").Body(app.Text("button")),
		app.Button().Class("button").Body(app.Text("button")),
		app.Button().Class("button is-primary is-outlined").Body(app.Text("button")),
		app.Button().Class("button is-loading").Body(app.Text("button")),
		app.Button().Class("button").Attr("disabled", true).Body(app.Text("button")),
	)
}
