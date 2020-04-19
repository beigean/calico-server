// Package auth ref. https://qiita.com/po3rin/items/740445d21487dfcb5d9f
package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

type jwtToken struct {
	Token string `json:"token"`
}

// GetTokenHandler ...
func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "54546557354"
	claims["name"] = "kazu"
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	var jsonToken jwtToken
	jsonToken.Token, _ = token.SignedString([]byte(os.Getenv("SIGNINKEY")))

	json.NewEncoder(w).Encode(jsonToken)
}

// JwtMiddleware ...
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINKEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

var AuthTest = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok!!\n"))
})
