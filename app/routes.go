package app

import "github.com/Basu008/Better-ESPN/api"

func (app *App) setRoutes() {
	app.Post("/player", app.handleRequest(api.CreatePlayer))
}
