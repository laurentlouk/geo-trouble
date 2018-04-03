package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Represents an accident, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Accident struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Date   time.Time     `bson:"date" json:"date"`
	City   string        `bson:"city" json:"city"`
	Number int           `bson:"number" json:"number"`
}
