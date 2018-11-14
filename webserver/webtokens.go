package webserver

import (
	"log"
	"net/http"
)

// this is file is a middleware for authenticatio for users, users will be assigned a token
// when they are logged on and this token will be destroyed if they leave the page or
// logouts

type authentication struct {
	tokens map[string]string
}

func (a *authentication) Middleware(next http.Handler) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r http.Request) {
		token := r.Header.Get("Announcement-Token")

		if user, ok := a.tokens[token]; ok {
			// We have now found a token from a user we can now let him
			// pass for now, and logs to console that this person is logged in
			log.Printf("User has been authorized %s\n", user)
			// Pass the request to next handler
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		}

	})
}
