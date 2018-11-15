package webserver

import (
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

// ServerStartup returns a bool to main to locate if the server start up
// started up like expected and returns false if it happend anything unusually
func ServerStartup() {
	port := os.Getenv("PORT")

	r := mux.NewRouter()

	r.HandleFunc("/Home", MainPage)
	r.HandleFunc("/Announce", Routing)

	//	authMiddleware := authentication{}

	log.Fatal(http.ListenAndServe(":"+port, r))

}
