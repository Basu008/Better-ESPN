package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	UserName string             `json:"username" bson:"user_name"`
	Passowrd string             `json:"password" bson:"password"`
}
