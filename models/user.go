package models

import "time"

type User struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
