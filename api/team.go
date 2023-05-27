package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Basu008/Better-ESPN/helpers"
	"github.com/Basu008/Better-ESPN/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const teamsCollection string = "teams"

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
	team.Players = playerIds
	_, err := db.Collection(teamsCollection).InsertOne(context.Background(), team)
	if err != nil {
		CreateNewResponse(w, http.StatusInternalServerError, &Response{false, "Couldn't add data to the server", nil})
		return
	}
	CreateNewResponse(w, http.StatusAccepted, &Response{true, "", true})
}

func GetTeamById(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idFromQuery := r.URL.Query().Get("id")
	if idFromQuery == "" {
		GetAllTeams(db, w, r)
		return
	}
	id, err := primitive.ObjectIDFromHex(idFromQuery)
	if err != nil {
		CreateNewResponse(w, http.StatusBadRequest, &Response{false, "Invalid Id", nil})
		return
	}
	filter := bson.D{
		{Key: "_id", Value: id},
	}
	var team model.Team
	mongoErr := db.Collection(teamsCollection).FindOne(context.TODO(), filter).Decode(&team)
	if mongoErr != nil {
		switch mongoErr {
		case mongo.ErrNoDocuments:
			CreateNewResponse(w, http.StatusNotFound, &Response{false, "No team with this id exists", nil})
		default:
			CreateNewResponse(w, http.StatusInternalServerError, &Response{false, "Couldn't fetch documents", nil})
		}
		return
	}
	CreateNewResponse(w, http.StatusAccepted, &Response{true, "", team})
}

func GetAllTeams(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	var teams = []model.Team{}
	curr, err := db.Collection(teamsCollection).Find(context.TODO(), bson.D{})
	if err != nil {
		CreateNewResponse(w, http.StatusInternalServerError, &Response{false, "Error Fetching data", nil})
		return
	}
	defer func() {
		err := curr.Close(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
	}()
	cursorErr := curr.All(context.TODO(), &teams)
	if cursorErr != nil {
		CreateNewResponse(w, http.StatusInternalServerError, &Response{false, "Couldn't fetch documents", nil})
		return
	}
	CreateNewResponse(w, http.StatusAccepted, &Response{true, "", teams})

}
