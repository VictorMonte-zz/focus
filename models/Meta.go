package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	//Meta : represent clients meta
	Meta struct {
		ID     bson.ObjectId `json:"id" bson:"_id"`
		Start  time.Time     `json:"start" bson:"start"`
		Value  float32       `json:"value" bson:"value"`
		Months int           `json:"months" bson:"months"`
	}
)
