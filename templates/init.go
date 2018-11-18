package templates

import (
	"log"
	"net/http"
	"text/template"
)

// Index s
func Index(w http.ResponseWriter) {
	tpl, err := template.ParseFiles("index.html")

	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(w, nil)
}

// Login s
func Login(w http.ResponseWriter) {
	tpl, err := template.ParseFiles("login.html")

	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(w, nil)
}

// Announce g
func Announce(w http.ResponseWriter) {
	tpl, err := template.ParseFiles("index.html")

	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(w, nil)
}

// Register g
func Register(w http.ResponseWriter) {
	tpl, err := template.ParseFiles("index.html")

	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(w, nil)
}
