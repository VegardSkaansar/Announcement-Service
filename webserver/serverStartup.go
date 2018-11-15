package webserver

import (
	"goprojects/AnnonceService/database"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

var tpl *template.Template

// Init initialise html directory
func Init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(wd)
	tpl = template.Must(template.ParseGlob(parent + "/src/goprojects/AnnonceService/html/*.html"))
}

// ServerRequest takes care of the routing
// and sends the user to right place
func ServerRequest() {
	port := os.Getenv("PORT")

	r := mux.NewRouter()
	r.HandleFunc("/", (func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, "/Home", 301) }))
	r.HandleFunc("/Home", MainPage)
	r.HandleFunc("/Announce", Routing)

	log.Fatal(http.ListenAndServe(":"+port, r))

}

// ServerStart initialize our global db so we can use it to
// interact with mlabs, and the interface depends on which user you
// are in this case admin or user
func ServerStart() {
	database.GlobalDBAdmin = &database.MongoDB{
		"mongodb://admin:admin123@ds039311.mlab.com:39311/announce",
		"announce",
		"user",
	}

	database.GlobalDBUser = &database.MongoDB{
		"mongodb://admin:admin123@ds039311.mlab.com:39311/announce",
		"announce",
		"user",
	}

	database.GlobalDBUser.Init()
	database.GlobalDBAdmin.Init()
}
