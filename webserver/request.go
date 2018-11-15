package webserver

import (
	"net/http"
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

}

// MainPage Displays this page if your not logged in
func MainPage(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "index.html", nil)

}
