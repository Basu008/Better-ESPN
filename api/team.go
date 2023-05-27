package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Basu008/Better-ESPN/helpers"
	"github.com/Basu008/Better-ESPN/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collection string = "teams"

func CreateTeam(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var teamRequest helpers.Team
	var team model.Team
	decodingErr := json.NewDecoder(r.Body).Decode(&teamRequest)
	if decodingErr != nil {
		CreateNewResponse(w, http.StatusBadRequest, &Response{false, "Issue with JSON Body", nil})
		return
	}
	team.Name = teamRequest.Name
	var playerIds = []primitive.ObjectID{}
	for _, playerId := range teamRequest.Players {
		idInBsonFormat, err := primitive.ObjectIDFromHex(playerId)
		if err != nil {
			CreateNewResponse(w, http.StatusBadRequest, &Response{false, "Invalid Ids", nil})
			return
		}
		playerIds = append(playerIds, idInBsonFormat)
	}
	if len(playerIds) <= 1 {
		CreateNewResponse(w, http.StatusBadRequest, &Response{false, "Add at least 2 players to the team :)", nil})
		return
	}
	_, err := db.Collection(collection).InsertOne(context.Background(), team)
	if err != nil {
		CreateNewResponse(w, http.StatusInternalServerError, &Response{false, "Couldn't add data to the server", nil})
		return
	}
	CreateNewResponse(w, http.StatusAccepted, &Response{true, "", true})
}
