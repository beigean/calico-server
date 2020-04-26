// Package auth ref. https://qiita.com/po3rin/items/740445d21487dfcb5d9f
package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type JwtToken struct {
	Token string `json:"token"`
}

type PrivateClaims struct {
	UserID    int    `json:"calico/user-id"`
	Mail      string `json:"calico/user-mail"`
	Name      string `json:"calico/user-name"`
	CreatedAt string `json:"calico/user-created_at"`
}

// MyClaims ...
type MyClaims struct {
	PrivateClaims
	jwt.StandardClaims
}

// GetTokenHandler ...
func CreateJwt(privateClaims *PrivateClaims) *JwtToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &MyClaims{
		*privateClaims,
		jwt.StandardClaims{
			Audience:  "localhost",
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Id:        "test",
			IssuedAt:  time.Now().Unix(),
			Issuer:    "localhost",
			NotBefore: time.Now().Add(time.Second * 5).Unix(),
			Subject:   "AccessToken",
		},
	})

	var jsonToken JwtToken
	jsonToken.Token, _ = token.SignedString(getSecretKey())

	return &jsonToken
}

// JwtMiddleware ...
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return getSecretKey(), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

// OnlyPersonMiddleware ...
func OnlyPersonMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := jwtmiddleware.FromAuthHeader(r)
		if err != nil {
			log.Fatal(err)
		}

		var claims MyClaims
		err = claims.GetFromTokenString(tokenString)
		if err != nil {
			log.Fatal(err)
		}

		if reqUserID, _ := strconv.Atoi(mux.Vars(r)["id"]); reqUserID != claims.UserID {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func getSecretKey() []byte {
	return []byte(os.Getenv("SIGNINKEY"))
}

// GetFromTokenString ...
func (c *MyClaims) GetFromTokenString(tokenString string) error {
	var claims MyClaims
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return getSecretKey(), nil
	})
	if err != nil {
		log.Fatal(err)
	}
	c = &claims
	return nil
}

// AuthTest ...
var AuthTest = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	headerAuthorization := strings.Split(r.Header.Get("Authorization"), " ")

	// check "Bearer"

	var claims MyClaims
	token, _ := jwt.ParseWithClaims(headerAuthorization[1], &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINKEY")), nil
	})
	fmt.Println(token.Valid)

	json.NewEncoder(w).Encode(claims)
})
