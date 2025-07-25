package models

import "time"

type Event struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Date        string    `json:"date" bson:"date"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
