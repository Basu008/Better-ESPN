package api

import (
	"encoding/json"
	"net/http"

	"github.com/Basu008/Better-ESPN/helpers"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user helpers.Login
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		CreateNewResponse(w, http.StatusBadRequest, &Response{false, "Issue with JSON Body", nil})
		return
	}
	CreateNewResponse(w, http.StatusAccepted, &Response{true, "", user})
}
