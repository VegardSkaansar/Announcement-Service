package webserver

import (
	"goprojects/AnnonceService/database"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var tpl *template.Template

// Init initialise html directory
func Init() {

	tpl = template.Must(template.ParseGlob("templates"))
}

// ServerRequest takes care of the routing
// and sends the user to right place
func ServerRequest() {
	port := os.Getenv("PORT")

	r := mux.NewRouter()
	r.HandleFunc("/", (func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, "/Home", 301) }))
	r.HandleFunc("/home", MainPage)
	r.Handle("/announce", isAuthorized(Routing))
	r.HandleFunc("/login", Login)
	r.HandleFunc("/register", Register)

	log.Fatal(http.ListenAndServe(":"+port, r))

}

// ServerStart initialize our global db so we can use it to
// interact with mlabs, and the interface depends on which user you
// are in this case admin or user
func ServerStart() {
	database.GlobalDBAdmin = &database.MongoDB{
		DatabaseURL:      "mongodb://admin:admin123@ds039311.mlab.com:39311/announce",
		DatabaseName:     "announce",
		DatabaseAnnounce: "user",
	}

	database.GlobalDBUser = &database.MongoDB{
		DatabaseURL:      "mongodb://admin:admin123@ds039311.mlab.com:39311/announce",
		DatabaseName:     "announce",
		DatabaseAnnounce: "user",
	}

	database.GlobalDBUser.Init()
	database.GlobalDBAdmin.Init()
}
