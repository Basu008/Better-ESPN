package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Basu008/Better-ESPN/helpers"
	"github.com/Basu008/Better-ESPN/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const userCollection string = "user"

func SignUp(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	var signUpBody helpers.SignUp
	if err := json.NewDecoder(r.Body).Decode(&signUpBody); err != nil {
		CreateNewResponse(w, http.StatusBadRequest, &Response{false, "Issue with JSON Body", nil})
		return
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpBody.Passowrd), 4)
	if err != nil {
		CreateNewResponse(w, http.StatusNotAcceptable, &Response{false, "Something's wrong", nil})
		return
	}
	user.Name = signUpBody.Name
	user.UserName = signUpBody.Username
	user.Passowrd = string(encryptedPassword)
	res, mongoErr := db.Collection(userCollection).InsertOne(context.Background(), user)
	if mongoErr != nil {
		CreateNewResponse(w, http.StatusInternalServerError, &Response{false, "Can't add data to the db", nil})
		return
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	// claim := auth.UserClaim{
	// 	Id:       user.ID.Hex(),
	// 	Name:     user.Name,
	// 	Username: user.UserName,
	// }
	// token, _ := claim.SignAuthToken()
	CreateNewResponse(w, http.StatusAccepted, &Response{true, "", true})
}
