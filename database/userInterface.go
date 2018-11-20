package database

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GlobalDBUser both admin and normal user should be able to interact with this page
// depending on their role, a admin can modify which user announcement the admins wants.
// While a user with no admin role can only interact on the users own announcement
var GlobalDBUser StorageAnnouncement

// StorageAnnouncement non admin inteface and here can clients
// add their Annoucements delete or modify
type StorageAnnouncement interface {
	Init()
	AddAnnouncement(ad Announce) bool
	DeleteAnnouncement(title string) bool
	UserAnnouncement(ID bson.ObjectId) []Announce
}

// Init test if the connection happend
// uses init in Admin userface to but are not declared there
func (db *MongoDB) Init() {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

}

// AddAnnouncement adds a users announce
func (db *MongoDB) AddAnnouncement(ad Announce) bool {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.DatabaseName).C(db.DatabaseAnnounce).Insert(ad)

	if err != nil {
		return false
	}
	return true
}

// DeleteAnnouncement takes a announce with a title and deletes it
func (db *MongoDB) DeleteAnnouncement(title string) bool {

	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.DatabaseName).C(db.DatabaseName).Remove(bson.M{"title": title})

	if err != nil {
		return false
	}
	return true
}

// UserAnnouncement takes a announce with a title and deletes it
func (db *MongoDB) UserAnnouncement(ID bson.ObjectId) []Announce {

	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var ann []Announce
	err = session.DB(db.DatabaseName).C(db.DatabaseName).Find(bson.M{"_id": ID}).All(&ann)
	log.Println(ann)
	if err != nil {
		return []Announce{}
	}
	return ann
}
