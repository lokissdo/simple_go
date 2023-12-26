package models


type Product struct {
	ID          int    `json:"id" bson:"_id"`
	Image       string    `json:"image" bson:"image"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Price       float64   `json:"price" bson:"price"`
	Color       string    `json:"color" bson:"color"`
}

