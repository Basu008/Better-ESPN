package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Basu008/Better-ESPN/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collName = "player"

func CreatePlayer(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	//Set the heder meaning the type of data that is being used ie. JSON
	w.Header().Set("Content-Type", "application/json")
	//Then we will fetch the JSON body that the user must've provided
	var player model.Player
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		//This will let the user know about to misformed JSON body
		responseBody := Response{false, "Issue with JSON Body", nil}
		CreateNewResponse(w, http.StatusBadRequest, &responseBody)
		return
	}

	//To set the initial
	player.PlayerProfile.Stats = &model.Stats{
		Goals:  0,
		Assist: 0,
		Fouls:  0,
	}

	player.CreatedAt = time.Now().UTC()

	//Now, if there is no error, we will send the player data to the DB
	result, err := db.Collection(collName).InsertOne(context.Background(), player)

	if err != nil {
		CreateNewResponse(w, http.StatusInternalServerError, &Response{false, "Player can't be create. Try Again!", nil})
		return
	}

	//Finally, we create a response body for success state

	player.ID = result.InsertedID.(primitive.ObjectID)
	CreateNewResponse(w, http.StatusCreated, &Response{true, "", player})
}
