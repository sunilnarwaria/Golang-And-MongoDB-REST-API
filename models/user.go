package models

import "gopkg.in/mgo.v2/bson"

// User represents a user in the database
type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Age      int           `json:"age" bson:"age"`
	Location string        `json:"location" bson:"location"`
}
