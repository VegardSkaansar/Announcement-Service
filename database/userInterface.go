package database

import mgo "gopkg.in/mgo.v2"

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
	ModifyAnnouncement(title string) bool
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
