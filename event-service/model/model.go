package model

import "gopkg.in/mgo.v2/bson"

type Event struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name"`
	Duration  int64         `json:"duration"`
	StartDate int64         `json:"startDate"`
	EndDate   int64         `json:"endDate"`
	Location  Location      `json:"location"`
}

type Location struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name"`
	Address  string        `json:"address"`
	Country  string        `json:"country"`
	OpensAt  int           `json:"opensAt"`
	ClosesAt int           `json:"closesAt"`
	Hall     []Hall        `json:"availableIn"`
}

type Hall struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Capacity int    `json:"capacity"`
}
