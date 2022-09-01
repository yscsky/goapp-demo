package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type cityData struct {
	name        string
	description string
	image       string
}

var cityStore = map[string]cityData{
	"paris": {
		name:  "Paris",
		image: "/web/paris.jpg",
		description: `
Paris, France's capital, is a major European city and a global center for art,
fashion, gastronomy and culture. Its 19th-century cityscape is crisscrossed by
wide boulevards and the River Seine. Beyond such landmarks as the Eiffel Tower
and the 12th-century, Gothic Notre-Dame cathedral, the city is known for its
cafe culture and designer boutiques along the Rue du Faubourg Saint-Honoré.
			`,
	},
	"sf": {
		name:  "San Francisco",
		image: "/web/sf.jpg",
		description: `
San Francisco, officially City and County of San Francisco and colloquially
known as SF, San Fran or "The City", is a city in—and the cultural, commercial,
and financial center of—Northern California.
			`,
	},
	"beijing": {
		name:  "北京市",
		image: "/web/beijing.jpg",
		description: `
Beijing, China's sprawling capital, has history stretching back 3 millennia. Yet
it's known as much for modern architecture as its ancient sites such as the
grand Forbidden City complex, the imperial palace during the Ming and Qing
dynasties. Nearby, the massive Tiananmen Square pedestrian plaza is the site of
Mao Zedong's mausoleum and the National Museum of China, displaying a vast
collection of cultural relics.
			`,
	},
}

type city struct {
	app.Compo

	data cityData
}

func (c *city) OnNav(ctx app.Context) {
	key := ctx.Page().URL().Query().Get("city")
	if key == "" {
		key = "paris"
	}
	data, ok := cityStore[key]
	if !ok {
		ctx.Navigate("/notfound")
		return
	}
	c.data = data
	c.Update()
}

func (c *city) Render() app.UI {
	return app.Div().Body(
		app.Main().Class("city").Style("background-image", fmt.Sprintf("url('%s')", c.data.image)).Body(
			app.H1().Class("city-title").Body(app.Text(c.data.name)),
			app.P().Class("city-description").Body(app.Text(c.data.description)),
			app.P().Class("city-links").Body(
				app.A().Class("app-button").Href("/city?city=beijing").Body(app.Text(cityStore["beijing"].name)),
				app.A().Class("app-button").Href("/city?city=paris").Body(app.Text(cityStore["paris"].name)),
				app.A().Class("app-button").Href("/city?city=sf").Body(app.Text(cityStore["sf"].name)),
			),
		),
	)
}
