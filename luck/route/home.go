package route

import (
	"fmt"
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Home struct {
	app.Compo

	luckNumber int
	games      []game
}

func (h *Home) OnMount(ctx app.Context) {
	myGames = make(map[string]game, len(games))
	if err := ctx.LocalStorage().Get("myGames", &myGames); err != nil {
		fmt.Println("LocalStorage Get myGames err:", err)
	}
	for k := range myGames {
		if _, ok := games[k]; !ok {
			delete(myGames, k)
		}
	}
	for k, v := range games {
		if old, ok := myGames[k]; ok {
			v.Enabled = old.Enabled
		}
		myGames[k] = v
	}
}

func (h *Home) OnNav(ctx app.Context) {
	if err := ctx.LocalStorage().Get("luckNumber", &h.luckNumber); err != nil {
		fmt.Println("LocalStorage Get luckNumber err:", err)
	}
	games := gameList(myGames)
	sortGameList(games)
	h.games = games
	h.Update()
}

func (h *Home) Render() app.UI {
	return app.Div().Class("layout").Body(
		app.Div().Class("background").Body(
			app.Main().Body(
				app.H1().Body(app.Text("What is your lucky number?")),
				app.Input().Class("section").Type("number").Placeholder("Lucky number").
					Value(h.luckNumber).OnChange(h.onLuckNumberChange).OnKeyUp(h.onLuckNumberChange),
				app.Table().Class("section").Body(
					app.Range(h.games).Slice(func(i int) app.UI {
						game := h.games[i]
						if !game.Enabled {
							return nil
						}
						return app.Div().Body(
							app.Td().Body(
								app.Div().Body(app.Text(game.Name)),
								app.Div().Class("subtitle accent").Body(app.Text(game.Location)),
							),
							app.Td().Class("draw").Body(app.Text(game.draw(h.luckNumber))),
						)
					}),
				),
				app.A().Class("app-button section").Href("select").Body(app.Text("Select other games")),
			),
		),
	)
}

func (h *Home) onLuckNumberChange(ctx app.Context, e app.Event) {
	number := ctx.JSSrc().Get("value").String()
	h.luckNumber, _ = strconv.Atoi(number)
	h.Update()

	if err := ctx.LocalStorage().Set("luckNumber", h.luckNumber); err != nil {
		fmt.Println("LocalStorage Set luckNumber err:", err)
	}
}
