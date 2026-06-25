package models

import (
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"go.mongodb.org/mongo-driver/v2/bson"

)

// "time"


type User struct{
	ID bson.ObjectID `bson:"_ id,omitempty" json:"id`
	Username string  `bson:"username" json:"username"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password omitempty"`
	Avatar string `bson:"avatar" json:"avatar"`
	CreatedAt time.Time `bson:"created at" json:"created at"`
	UpdatedAt time.Time `bson:"updated at" json:"updated at"`
}

type RegisterRequest struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct{
	ID string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}