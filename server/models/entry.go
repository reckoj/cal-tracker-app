package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct{
	Id  		primitive.ObjectID `bson: "id"`	
	Dish		*string 		   `json:"dish"`
	Fat			*float64		   `json:"fat"`
	Ingridients	*string			   `json:"ingredieints"`
	Calories	*string			   `json:"calories"`
} 