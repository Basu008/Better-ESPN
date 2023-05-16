package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	PlayerProfile *PlayerProfile     `json:"player_profile,omitempty" bson:"player_profile,omitempty"`
	CreatedAt     time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type PlayerProfile struct {
	Grade string `json:"grade,omitempty" bson:"grade,omitempty"`
	Stats *Stats `json:"stats,omitempty" bson:"stats,omitempty"`
}

type Stats struct {
	Goals    int `json:"goals,omitempty" bson:"goals,omitempty"`
	Baskets  int `json:"baskets,omitempty" bson:"baskets,omitempty"`
	Assist   int `json:"assists,omitempty" bson:"assists,omitempty"`
	Fouls    int `json:"fouls,omitempty" bson:"fouls,omitempty"`
	Rebounds int `json:"rebounds,omitempty" bson:"rebounds,omitempty"`
}
