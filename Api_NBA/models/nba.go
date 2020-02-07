package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Nba struct {
	ID             bson.ObjectId `bson:"_id" json:"id"`
	Year           int           `bson:"Year" json:"Year"`
	Player         string        `bson:"Player" json:"Player"`
	Pos            string        `bson:"Pos" json:"Pos"`
	HT             string        `bson:"HT" json:"HT"`
	WT             int           `bson:"WT" json:"WT"`
	Team           string        `bson:"Team" json:"Team"`
	SelectionType  string        `bson:"Selection Type" json:"Selection Type"`
	NBADraftStatus string        `bson:"NBA Draft Status" json:"NBA Draft Status"`
	Nationality    string        `bson:"Nationality" json:"Nationality"`
}

// ID	Year	Player	Pos	HT	WT	Team	Selection Type	NBA Draft Status	Nationality
