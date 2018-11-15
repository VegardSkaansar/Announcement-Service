package database

import "gopkg.in/mgo.v2/bson"

// MongoDB is the server connection to mlabs with mongodb
// a non sql database connection
type MongoDB struct {
	DatabaseURL      string
	DatabaseName     string
	DatabaseUser     string
	DatabaseAnnounce string
}

// User is the db structure of all the users in the server
// Role and password will be stored in db but no other way to get
// this information from the db
type User struct {
	ObjectID  bson.ObjectId `json:"-" bson:"_id"`
	Username  string        `json:"username" bson:"username"`
	Password  string        `json:"-" bson:"password"`
	FirstName string        `json:"firstname"`
	LastName  string        `json:"lastname"`
	Year      string        `json:"year"`
	Role      string        `json:"-" bson:"role"`
	anno      Announce
}

// Announce is the db structure for all the announcement given from
// a user
type Announce struct {
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	ObjectID    bson.ObjectId `json:"-" bson:"_id"`
	Cost        int           `json:"cost"`
}

// StorageUser interface for the database for adding users and deleting
// Only admins can use this interface
type StorageUser interface {
	AddUser(person User)
	DeleteUser(username string)
	ExistUser(username string) bool
}

// AddUser adds a user to the datastructure
func (db *MongoDB) AddUser(person User) {

}

// DeleteUser deletes a user from the db
func (db *MongoDB) DeleteUser(username string) {

}

// ExistUser checks if a user with this username exist
func (db *MongoDB) ExistUser(username string) bool {

	return true
}

// StorageAnnouncement non admin inteface and here can clients
// add their Annoucements delete or modify
type StorageAnnouncement interface {
	AddAnnouncement(ad Announce) bool
	DeleteAnnouncement(title string) bool
	ModifyAnnouncement(title string) bool
}

// AddAnnouncement adds a users announce
func (db *MongoDB) AddAnnouncement(ad Announce) bool {

	return true
}

// DeleteAnnouncement takes a announce with a title and deletes it
func (db *MongoDB) DeleteAnnouncement(title string) bool {

	return true
}

// ModifyAnnouncement Takes a title and modifies what information the user wants
func (db *MongoDB) ModifyAnnouncement(title string) bool {

	return true
}
