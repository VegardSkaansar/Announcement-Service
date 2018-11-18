package webserver

import (
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
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
	log.Println(pwd)
	log.Println(hashed)
	log.Println(hashAndSalt(pwd))
	byteHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println("password true")
	return true
}

// hash keys and blockKey should not been uploaded to git or anywhere,
// but this task it doesnt really matter for us, and we will show others
// what we have done.
// this keys should be 32-64 bit
var hashKey = []byte("Something-secret")
var s = securecookie.New(hashKey, nil)

// SetCookie sets a token cookie thats encoded
func SetCookie(token string, w http.ResponseWriter, r *http.Request) {
	encoded, err := s.Encode("Authorization", token)
	if err == nil {
		cookie := &http.Cookie{
			Name:       "Authorization",
			Value:      encoded,
			RawExpires: "0",
		}
		http.SetCookie(w, cookie)
	}
}

// ReadCookie reads a cookie and returns the token
func ReadCookie(w http.ResponseWriter, r *http.Request) string {

	cookie, err := r.Cookie("Authorization")

	if err != nil {
		http.Error(w, "Not authorization", http.StatusForbidden)
		return ""
	}
	var value string
	err = s.Decode("Authorization", cookie.Value, &value)
	if err != nil {
		http.Error(w, "Not authorization", http.StatusForbidden)
		return ""
	}
	return value
}
