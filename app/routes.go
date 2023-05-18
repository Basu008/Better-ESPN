package app

import "github.com/Basu008/Better-ESPN/api"

func (app *App) setRoutes() {
	app.Post("/player", app.handleRequest(api.CreatePlayer))
	app.Get("/players", app.handleRequest(api.GetAllPlayers))
	app.Get("/player/{id}", app.handleRequest(api.GetPlayerById))
}
