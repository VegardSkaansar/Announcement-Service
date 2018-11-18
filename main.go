package main

import (
	"goprojects/AnnonceService/webserver"
)

// This is the main and are supposed to delegate the work to the other packages
// This can be comparable to a controller in model-view-controller. We can think
// of the packages as the model and html packes as the view where the customers
// can get the information they need

// main is in control of starting the server, assign a port and ipadress to the server
// main will listen to the ipadress/port for request and send it to the right
// handler
func main() {
	webserver.ServerStart()
	webserver.ServerRequest()
}
