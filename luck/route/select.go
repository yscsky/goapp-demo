package route

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type GameSelect struct {
	app.Compo

	filter string
	games  []game
}

func (s *GameSelect) OnNav(ctx app.Context) {
	s.showGames()
}

func (s *GameSelect) Render() app.UI {
	return app.Div().Class("layout").Body(
		app.Div().Class("background").Body(
			app.Main().Body(
				app.H1().Body(app.Text("Select games")),
				app.Input().Class("section").Type("text").Placeholder("Filter").
					Value(s.filter).OnChange(s.onFilterChange).OnKeyUp(s.onFilterChange),
				app.Table().Class("section").Body(
					app.Range(s.games).Slice(func(i int) app.UI {
						game := s.games[i]
						return app.Tr().Class("gamerow").DataSet("id", game.ID).
							OnClick(s.onGameSelect).Body(
							app.Td().Body(
								app.Div().Class("title").Body(app.Text(game.Name)),
								app.Div().Class("subtitle").Body(app.Text(game.Location)),
							),
							app.Td().Body(
								app.If(game.Enabled, app.Div().Class("selected").Body(app.Text("✓"))),
							),
						)
					}),
				),
				app.A().Class("app-button section").Href("/").Body(app.Text("Done")),
			),
		),
	)
}

func (s *GameSelect) showGames() {
	games := gameList(myGames)
	games = filterGameList(games, s.filter)
	sortGameList(games)
	s.games = games
	s.Update()
}

func (s *GameSelect) onFilterChange(ctx app.Context, e app.Event) {
	filter := ctx.JSSrc().Get("value").String()
	s.filter = filter
	s.Update()
	s.showGames()
}

func (s *GameSelect) onGameSelect(ctx app.Context, e app.Event) {
	id := ctx.JSSrc().Get("dataset").Get("id").String()
	game := myGames[id]
	game.Enabled = !game.Enabled
	myGames[id] = game

	if err := ctx.LocalStorage().Set("myGames", myGames); err != nil {
		fmt.Println("LocalStorage Set myGames err:", err)
	}

	s.Update()
	s.showGames()
}
