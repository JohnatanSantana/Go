package dao

import (
	"log"

	. "github.com/johnatansantana/Go/Api_NBA/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type NbaDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "nba"
)

func (m *NbaDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *NbaDAO) GetAll() ([]Nba, error) {
	var nbas []Nba
	err := db.C(COLLECTION).Find(bson.M{}).All(&nbas)
	return nbas, err
}

func (m *NbaDAO) GetByID(id string) (Nba, error) {
	var nba Nba
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&nba)
	return nba, err
}

func (m *NbaDAO) Create(nba Nba) error {
	err := db.C(COLLECTION).Insert(&nba)
	return err
}

func (m *NbaDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *NbaDAO) Update(id string, nba Nba) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &nba)
	return err
}
