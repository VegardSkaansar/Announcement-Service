package webserver

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// hashAndSalt will use a hashing algorithim with our password and then give it some
// salt so it will be harder to crack
func hashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// comparePassword compares the login password up to the hashed password
// if the hashed string is the same as the password we return true,
// otherwise we will get false
func comparePassword(pwd string, hashed string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
