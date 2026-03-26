package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type TodoStatus string

const (
	StatusComplete   TodoStatus = "complete"
	StatusIncomplete TodoStatus = "incomplete"
)

type Todo struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Status      TodoStatus         `json:"status" bson:"status"`
}