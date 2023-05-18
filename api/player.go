package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Basu008/Better-ESPN/helpers"
	"github.com/Basu008/Better-ESPN/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collName = "player"

func CreatePlayer(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	//Set the heder meaning the type of data that is being used ie. JSON
	w.Header().Set("Content-Type", "application/json")
	//Then we will fetch the JSON body that the user must've provided
	var player model.Player
	var requestBody helpers.CreatePlayerRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		//This will let the user know about to misformed JSON body
		CreateNewResponse(w, http.StatusBadRequest, &Response{false, "Issue with JSON Body", nil})
		return
	}
	if !requestBody.IsCreatePlayerRequestBodyValid() {
		CreateNewResponse(w, http.StatusBadRequest, &Response{false, "Inputs given are invalid", nil})
		return
	}
	//Now we set the data
	player.Name = requestBody.Name
	player.PlayerProfile = &model.PlayerProfile{
		Grade:    requestBody.Grade,
		Position: requestBody.Position,
		Stats: &model.Stats{
			Goals:  0,
			Assist: 0,
			Fouls:  0,
		},
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

func GetAllPlayers(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//First we get the cursor
	var players []model.Player
	cursor, err := db.Collection(collName).Find(context.TODO(), bson.M{})
	defer func() {
		err := cursor.Close(context.Background())
		if err != nil {
			log.Printf("Couldn't close cursor")
		}
	}()
	if err != nil {
		CreateNewResponse(w, http.StatusInternalServerError, &Response{false, "Coudln't find data", nil})
		return
	}
	cursorErr := cursor.All(context.Background(), &players)
	if cursorErr != nil {
		CreateNewResponse(w, http.StatusInternalServerError, &Response{false, "Coudln't find data", nil})
		return
	}
	CreateNewResponse(w, http.StatusAccepted, &Response{true, "", players})
}

func GetPlayerById(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//all the quries from the end point
	params := mux.Vars(r)
	//Then we get the id from the map of queries
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		CreateNewResponse(w, http.StatusBadRequest, &Response{false, "Invalid player ID", nil})
		return
	}
	//then we check in the db for the existance of document
	var player model.Player
	//How to get a document based on ID
	mongoErr := db.Collection(collName).FindOne(context.Background(), model.Player{ID: id}).Decode(&player)
	if mongoErr != nil {
		switch mongoErr {
		case mongo.ErrNoDocuments:
			CreateNewResponse(w, http.StatusNotFound, &Response{false, "No player with this id exists", nil})
		default:
			CreateNewResponse(w, http.StatusInternalServerError, &Response{false, "Couldn't fetch documents", nil})
		}
		return
	}
	CreateNewResponse(w, http.StatusAccepted, &Response{true, "", player})

}

func DeletePlayer(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Here we will only delete one player
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		CreateNewResponse(w, http.StatusBadRequest, &Response{false, "Invalid player ID", nil})
		return
	}
	_, deletionError := db.Collection(collName).DeleteOne(context.Background(), model.Player{ID: id})
	if deletionError != nil {
		switch deletionError {
		case mongo.ErrNoDocuments:
			CreateNewResponse(w, http.StatusNotFound, &Response{false, "No player with this id exists", nil})
		default:
			CreateNewResponse(w, http.StatusInternalServerError, &Response{false, "Couldn't fetch documents", nil})
		}
		return
	}
	CreateNewResponse(w, http.StatusAccepted, &Response{true, "", nil})

}
