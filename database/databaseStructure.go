package database

import "gopkg.in/mgo.v2/bson"

// MongoDB is the server connection to mlabs with mongodb
// a non sql database connection
type MongoDB struct {
	DatabaseURL      string
	DatabaseName     string
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
	Admin     bool          `json:"-" bson:"admin"`
}

// Announce is the db structure for all the announcement given from
// a user
type Announce struct {
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	ObjectID    bson.ObjectId `json:"-" bson:"_id"`
	Cost        int           `json:"cost"`
}

// Collection is how we will store our data in the db
// so we will have a user with its information and then
// have all of the announces this user have posted
type Collection struct {
	person User
	ads    []Announce
}
