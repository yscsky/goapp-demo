package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type hello struct {
	app.Compo

	Name string
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		// app.Div().Class("menu-button").OnClick(h.onMenuClick).Body(app.Text("☰")),
		app.Nav().Class("menu-button").Body(
			app.A().OnClick(h.onReload).Body(app.Text("Reload")),
			app.Text("|"),
			app.A().Style("color", "white").Href("/city").Body(app.Text("City demo")),
			app.Text("|"),
			app.A().Style("color", "white").Href("https://github.com/maxence-charriere/go-app").
				Body(app.Text("Go to repository")),
			app.Text("|"),
			app.A().Style("color", "white").Href("https://github.com/maxence-charriere/go-app-demo/tree/v6/demo").
				Body(app.Text("Demo sources")),
		),
		app.Main().Class("hello").Body(
			app.H1().Class("hello-title").Body(
				app.Text("Hello, "),
				app.If(h.Name != "", app.Text(h.Name)).Else(app.Text("World")),
			),
			app.Input().Class("hello-input").Value(h.Name).
				Placeholder("What is your name?").AutoFocus(true).OnChange(h.onInputChange),
		),
	)
}

func (h *hello) onInputChange(ctx app.Context, e app.Event) {
	h.Name = ctx.JSSrc().Get("value").String()
	h.Update()
}

func (h *hello) onReload(ctx app.Context, e app.Event) {
	ctx.Reload()
}
