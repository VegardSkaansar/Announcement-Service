package main

import (
	"fmt"
	"goprojects/AnnonceService/webserver"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

// This is the main and are supposed to delegate the work to the other packages
// This can be comparable to a controller in model-view-controller. We can think
// of the packages as the model and html packes as the view where the customers
// can get the information they need

// main is in control of starting the server, assign a port and ipadress to the server
// main will listen to the ipadress/port for request and send it to the right
// handler
func main() {
	Init()
	webserver.ServerStart()
	webserver.ServerRequest()
}

var tpl *template.Template

// Init initialise html directory
func Init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
