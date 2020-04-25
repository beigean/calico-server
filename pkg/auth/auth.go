// Package auth ref. https://qiita.com/po3rin/items/740445d21487dfcb5d9f
package auth

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/noobs9/calico-server/pkg/controller"
	"golang.org/x/crypto/bcrypt"
)

type jwtToken struct {
	Token string `json:"token"`
}

// GetTokenHandler ...
func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	len, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := make([]byte, len)
	len, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var bufReq controller.UserCols
	err = json.Unmarshal(body, &bufReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if bufReq.Mail == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if bufReq.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var bufDB controller.UserCols
	db, err := sqlx.Open(controller.KindDb, controller.Dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.QueryRowx("SELECT id, mail, password, name, created_at FROM users WHERE mail=?", bufReq.Mail).StructScan(&bufDB)
	if err == nil {
		// pass
	} else if err == sql.ErrNoRows {
		// email adress is not matched
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		log.Fatal(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(bufDB.Password), []byte(bufReq.Password))
	if err == nil {
		// pass
	} else {
		// password is not matched
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// fmt.Println(bufDB)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	// Registerd Claim
	claims["jti"] = "test"
	claims["iss"] = "localhost"
	claims["sub"] = "AccessToken"
	claims["aud"] = []string{"localhost"}
	claims["iat"] = time.Now()
	claims["npf"] = time.Now().Add(time.Second * 5).Unix()
	claims["exp"] = time.Now().Add(time.Minute).Unix()
	// Private Claim
	privatePrefix := "localhost"
	claims[privatePrefix+"id"] = bufDB.ID
	claims[privatePrefix+"mail"] = bufDB.Mail
	claims[privatePrefix+"name"] = bufDB.Name
	claims[privatePrefix+"created_at"] = bufDB.CreatedAt

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

// AuthTest ...
var AuthTest = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok!!\n"))
})
