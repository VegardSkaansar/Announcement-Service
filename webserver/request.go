package webserver

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/VegardSkaansar/Announcement-Service/database"

	"gopkg.in/mgo.v2/bson"
)

// url should look like
// username/password Encrypted/ = root
// root/api should handdle all get request for announces
// root/api/new_announcement should be able to post a new announcement

// new user will be made through the html pages through input fields.
// we need to check for a new users username are already in use or not
// and check if the epost is already in use or not. No reasons for
// having two users on the same epost.
// this will be taken up from a webserver and add the user to the db
// if everything was okay

// only the user that added the announcement can delete it or modify it,
// And of course an admin can remove or modify announcements.

// Routing function will use regex, and redirect or check the urlpath
// and send the request to right handler
func Routing(w http.ResponseWriter, r *http.Request) {
	//CookieValue := ReadCookie(w, r)
	//username := decodeToken(CookieValue)
	ExecuteHTML(w, "templates/announce.html", nil)
	log.Println("given access to resources")

}

// MainPage Displays this page if your not logged in
func MainPage(w http.ResponseWriter, r *http.Request) {

	ExecuteHTML(w, "templates/index.html", nil)
}

// Login this will execute the login page for now
func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		//templates.Login(w)
		ExecuteHTML(w, "AbuHtml/noasLogin.html", nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		if database.GlobalDBAdmin.ExistUser(r.Form["username"][0]) && comparePassword(r.Form["password"][0], database.GlobalDBAdmin.GetUserPassword(r.Form["username"][0])) {

			token, err := createToken(r.Form["username"][0])
			if err != nil {
				http.Error(w, "Error: Internal Server Error", 500)
			}
			SetCookie(token.Token, w, r)
			JSONResponse(token, w)

		} else {
			http.Error(w, "This user does not exist", 401)
			return
		}

	} else {
		http.Error(w, http.StatusText(405), 405)
		return
	}

}

// Register this html page add a user page
func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		ExecuteHTML(w, "templates/register.html", nil)
	} else if r.Method == "POST" {
		r.ParseForm()

		if database.GlobalDBAdmin.ExistUser(r.Form["username"][0]) {
			http.Error(w, "This Username is already in use", 409)
			return
		}
		log.Println("New user")
		log.Println("username:", r.Form["username"])
		log.Println("password:", r.Form["password"])
		log.Println("first name:", r.Form["firstName"])
		log.Println("last name:", r.Form["lastName"])
		log.Println("Birth:", r.Form["year"])

		if r.Form["password"][0] != r.Form["confirm_password"][0] {
			ExecuteHTML(w, "templates/register.html", "password doesnt match")
			return
		}
		newUser := database.User{
			Username:  r.Form["username"][0],
			Password:  hashAndSalt(r.Form["password"][0]),
			FirstName: r.Form["firstName"][0],
			LastName:  r.Form["lastName"][0],
			Year:      r.Form["year"][0],
			Admin:     false,
		}

		var ann []database.Announce
		collection := database.Collection{
			ObjectID: bson.NewObjectId(),
			Person:   newUser,
			Ads:      ann,
		}
		database.GlobalDBAdmin.AddUser(collection)
	}
}

// JSONResponse a helper function
func JSONResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// MyAds you here from this page modify or delete or add your annouces
func MyAds(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		CookieValue := ReadCookie(w, r)
		username := decodeToken(CookieValue)

		array := database.GlobalDBAdmin.GetUser(username.(string))

		ExecuteHTML(w, "AbuHtml/myAds.html", array)

	}
}

// ExecuteHTML parising a template and execute it
func ExecuteHTML(w http.ResponseWriter, path string, message interface{}) {
	tpl, err := template.ParseFiles(path)

	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(w, message)
}
