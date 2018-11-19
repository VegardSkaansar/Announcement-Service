package database

import (
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
	AddAnnouncement(ad Announce, username string) bool
	DeleteAnnouncement(title string, username string) bool
	ModifyAnnouncement(ads Announce, username string) bool
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
func (db *MongoDB) AddAnnouncement(ad Announce, username string) bool {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	var result []Collection

	err = session.DB(db.DatabaseName).C(db.DatabaseAnnounce).Find(bson.M{}).All(&result)

	if err != nil {
		return false
	}
	for _, data := range result {
		if data.Username == username {
			data.Ads = append(data.Ads, ad)
			GlobalDBAdmin.DeleteUser(username)
			GlobalDBAdmin.AddUser(data)
			return true
		}
	}
	return false
}

// DeleteAnnouncement takes a announce with a title and deletes it
func (db *MongoDB) DeleteAnnouncement(title string, username string) bool {

	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	change := bson.M{"ads": bson.M{"$pull": bson.M{"title": title}}}
	match := bson.M{"person": bson.M{"$elemMatch": bson.M{"username": username}}}

	err = session.DB(db.DatabaseName).C(db.DatabaseName).Update(match, change)

	if err != nil {
		return false
	}
	return true
}

// ModifyAnnouncement Takes a title and modifies what information the user wants
// doesnt work yet will work when discussed
func (db *MongoDB) ModifyAnnouncement(ads Announce, username string) bool {

	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	change := bson.M{"ads": bson.M{"$set": bson.M{"title": ads.Title}}}
	match := bson.M{"person": bson.M{"$elemMatch": bson.M{"username": username}}}

	err = session.DB(db.DatabaseName).C(db.DatabaseName).Update(match, change)

	if err != nil {
		return false
	}
	return true
}
