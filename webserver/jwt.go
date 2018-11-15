package webserver

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// secret signing value are made here, security reasons should this have been a value that no one could see
// but because of this project we do it very simple and make it here
var hmacSecret = []byte("Secret")

// createToken takes a username and a password and uses hmac encoding and encodes it with the secret word
func createToken(username string, cryptPass string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": cryptPass,
		"exp":      0,
	})

	tokenString, err := token.SignedString(hmacSecret)
	return tokenString, err
}

// decodeToken takes a token from the header of a client and decodes it
func decodeToken(token string) (interface{}, interface{}) {
	data, err := jwt.Parse(token, func(tokenString *jwt.Token) (interface{}, error) {

		if _, ok := tokenString.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", tokenString.Header["alg"])
		}
		return hmacSecret, nil
	})
	if err != nil {
		return false, err
	}

	if claims, ok := data.Claims.(jwt.MapClaims); ok && data.Valid {
		return claims["username"], claims["password"]
	} else {
		return false, false
	}

}
