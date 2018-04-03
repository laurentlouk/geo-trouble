package dao

import (
	"log"
	"time"

	"github.com/laurentlouk/geo-trouble/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// AccidentsDAO credetials for db
type AccidentsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	// COLLECTION accidents
	COLLECTION = "accidents"
)

// Connect : Establish an connection to database
func (a *AccidentsDAO) Connect() {
	session, err := mgo.Dial(a.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(a.Database)
}

// FindByDateCity : Find with Date range
func (a *AccidentsDAO) FindByDateCity(start, end time.Time, city string) ([]models.Accident, error) {
	var accidents []models.Accident
	err := db.C(COLLECTION).Find(
		bson.M{
			"city": bson.M{"$eq": city},
			"date": bson.M{
				"$gte": start,
				"$lte": end,
			},
		}).All(&accidents)
	return accidents, err
}

// FindAll : Find list of accidents
func (a *AccidentsDAO) FindAll() ([]models.Accident, error) {
	var accidents []models.Accident
	err := db.C(COLLECTION).Find(bson.M{}).All(&accidents)
	return accidents, err
}

// FindByID : Find an accident by its id
func (a *AccidentsDAO) FindByID(id string) (models.Accident, error) {
	var accident models.Accident
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&accident)
	return accident, err
}

// Insert : Insert an accident into database
func (a *AccidentsDAO) Insert(accident models.Accident) error {
	err := db.C(COLLECTION).Insert(&accident)
	return err
}

// Update an existing movie
func (a *AccidentsDAO) Update(accident models.Accident) error {
	err := db.C(COLLECTION).UpdateId(accident.ID, &accident)
	return err
}

// Delete : Delete an existing movie
func (a *AccidentsDAO) Delete(accident models.Accident) error {
	err := db.C(COLLECTION).Remove(&accident)
	return err
}

// RemoveAll : remove all collection
func (a *AccidentsDAO) RemoveAll() {
	db.C(COLLECTION).RemoveAll(nil)
}
