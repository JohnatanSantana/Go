package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Year        string        `bson:"year" json:"year"`
	Player      string        `bson:"player" json:"player"`
	HT          string        `bson:"ht" json:"ht"`
	WT          string        `bson:"wt" json:"wt"`
	Team        string        `bson:"team" json:"team"`
	Selection   string        `bson:"selection" json:"selection"`
	Type        string        `bson:"type" json:"type"`
	NBA         string        `bson:"nba" json:"nba"`
	Draft       string        `bson:"draft" json:"draft"`
	Status      string        `bson:"status" json:"status"`
	Nationality string        `bson:"nationality" json:"nationality"`
}

// ID	Year	Player	Pos	HT	WT	Team	Selection Type	NBA Draft Status	Nationality
