package webserver

import (
	"fmt"
	"goprojects/AnnonceService/database"
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
func createToken(username string) JwtToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	tokenString, err := token.SignedString(hmacSecret)

	if err != nil {
		fmt.Println(err)
	}

	return JwtToken{tokenString}
}

// decodeToken takes a token from the header of a client and decodes it
func decodeToken(token string) string {
	data, err := jwt.Parse(token, func(tokenString *jwt.Token) (interface{}, error) {

		if _, ok := tokenString.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", tokenString.Header["alg"])
		}
		return hmacSecret, nil
	})
	if err != nil {
		return ""
	}

	if claims, ok := data.Claims.(jwt.MapClaims); ok && data.Valid {
		return claims["username"].(string)
	}
	return ""

}

// TokenMiddleware checks if the person are logged in or not
func TokenMiddleware(w http.ResponseWriter, r *http.Request) bool {

	// Get the token from Autherization header
	token := r.Header.Get("Authentication")
	jwtTokens := decodeToken(token)

	if database.GlobalDBAdmin.ExistUser(jwtTokens) {
		return true
	}
	// if this username doesnt exist either an deleted accounts token
	// or someone trying to get access to an account they shouldnt
	// get access to
	http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	return false

}
