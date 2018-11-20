package database

import (
	"log"

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
	GetUserPassword(username string) string
	GetUser(username string) bson.ObjectId
}

// AddUser adds a user to the datastructure
// with a empty annoucement array
func (db *MongoDB) AddUser(person User) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.DatabaseName).C(db.DatabaseAnnounce).Insert(person)

	if err != nil {
		log.Printf("Somethings wrong with Insert():%v", err.Error())
	}
}

// DeleteUser deletes a user from the db
func (db *MongoDB) DeleteUser(username string) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.DatabaseName).C(db.DatabaseAnnounce).Remove(bson.M{"username": username})

	if err != nil {
		log.Printf("Somethings wrong with Remove():%v", err.Error())
	}

}

// ExistUser checks if a user with this username exist
// people from other
func (db *MongoDB) ExistUser(username string) bool {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var result []User

	err = session.DB(db.DatabaseName).C(db.DatabaseAnnounce).Find(bson.M{}).All(&result)

	if err != nil {
		return false
	}
	for _, data := range result {
		if data.Username == username {
			return true
		}
	}
	return false

}

// GetUserPassword gets the password from a user else an error will be returned
func (db *MongoDB) GetUserPassword(username string) string {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	var pass []User

	err = session.DB(db.DatabaseName).C(db.DatabaseAnnounce).Find(bson.M{}).All(&pass)

	if err != nil {
		return ""
	}
	for _, data := range pass {
		if data.Username == username {
			return data.Password
		}
	}
	return ""
}

//GetUser gets the user
func (db *MongoDB) GetUser(username string) bson.ObjectId {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	var pass User

	err = session.DB(db.DatabaseName).C(db.DatabaseAnnounce).Find(bson.M{"username": username}).One(&pass)

	if err != nil {
		log.Println("failed to find object id")
	}

	return pass.ObjectID
}
