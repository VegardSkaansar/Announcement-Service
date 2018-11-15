package database

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GlobalDBAdmin is what a admin can do, and users with not admin shouldnt
// be able to interact with this interface
var GlobalDBAdmin StorageUser

// StorageUser interface for the database for adding users and deleting
// Only admins can use this interface
type StorageUser interface {
	Init()
	AddUser(person User)
	DeleteUser(username string)
	ExistUser(username string) bool
}

// AddUser adds a user to the datastructure
// with a empty annoucement array
func (db *MongoDB) AddUser(person Collection) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.DatabaseName).C(db.DatabaseAnnounce).Insert(person)

	if err != nil {
		fmt.Printf("Somethings wrong with Insert():%v", err.Error())
	}
}

// DeleteUser deletes a user from the db
func (db *MongoDB) DeleteUser(username string) {

}

// ExistUser checks if a user with this username exist
// people from other
func (db *MongoDB) ExistUser(username string) bool {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var result Collection

	err = session.DB(db.DatabaseName).C(db.DatabaseAnnounce).Find(bson.M{"person": bson.M{"$elemMatch": bson.M{"username": username}}}).One(&result)

	if err != nil {
		return false
	}
	return true
}
