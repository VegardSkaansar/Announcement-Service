package webserver

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// ServerStartup returns a bool to main to locate if the server start up
// started up like expected and returns false if it happend anything unusually
func ServerStartup() bool {
	port := os.Getenv("PORT")

	r := mux.NewRouter()
	r.HandleFunc("/Announce", Routing)

	authMiddleware := authentication{}

	log.Fatal(http.ListenAndServe(":"+port, r))

}
