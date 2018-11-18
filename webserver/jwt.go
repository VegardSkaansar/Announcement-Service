package webserver

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// secret signing value are made here, security reasons should this have been a value that no one could see
// but because of this project we do it very simple and make it here
var hmacSecret = []byte("Secret")

// JwtToken struct
type JwtToken struct {
	Token string `json:"token"`
}

// createToken takes a username and a password and uses hmac encoding and encodes it with the secret word
// this will only contain the username for several reasons. its bad to involve password even if it is encrypted
func createToken(username string) (JwtToken, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username

	tokenString, err := token.SignedString(hmacSecret)

	if err != nil {
		fmt.Errorf("Something didnt go well: %s", err.Error())
		return JwtToken{}, err
	}

	return JwtToken{tokenString}, nil
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CookieValue := ReadCookie(w, r)

		if CookieValue != "" {

			token, err := jwt.Parse(CookieValue, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return hmacSecret, nil
			})

			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			http.Error(w, "Not Authorized", http.StatusForbidden)
			return
		}
	})
}
