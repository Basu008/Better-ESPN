package app

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router   *mux.Router
	Database *mongo.Database
}
