package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type WatchList struct {
	/*Adaptando os campos da struct para o DB reconhecer o JSON*/
	// ID único
	ID     primitive.ObjectID `json:"_id, omitempty bson:_id,omitempty"`
	Anime  string             `json:"anime,omitempty"`
	Status bool               `json:"status,omitempty"`
	Nota   uint               `json:"nota,omitempty"`
}
