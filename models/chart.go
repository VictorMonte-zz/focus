package models

import "gopkg.in/mgo.v2/bson"

type (
	//Chart : represent data for chart
	Chart struct {
		ID           bson.ObjectId `json:"id" bson:"_id"`
		Labels       []string      `json:"labels" bson:"labels"`
		Custody      []float32     `json:"custody" bson:"custody"`
		Accomplished []float32     `json:"accomplished" bson:"accomplished"`
		Meta         []float32     `json:"meta" bson:"meta"`
	}
)
