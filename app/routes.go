package app

import "github.com/Basu008/Better-ESPN/api"

func (app *App) setRoutes() {
	app.Post("/player", app.handleRequest(api.CreatePlayer))
	app.Get("/players", app.handleRequest(api.GetAllPlayers))
	app.Get("/player/{id}", app.handleRequest(api.GetPlayerById))
	app.Get("/top-scorers", app.handleRequest(api.GetTopScorers))
	app.Delete("/player/{id}", app.handleRequest(api.DeletePlayer))
	app.Put("/player/{id}", app.handleRequest(api.UpdatePlayer))

	app.Post("/team", app.handleRequest(api.CreateTeam))
	app.Get("/team", app.handleRequest(api.GetTeamById))
}
