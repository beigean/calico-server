// Package auth ref. https://qiita.com/po3rin/items/740445d21487dfcb5d9f
package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/noobs9/calico-server/pkg/controller"
	"golang.org/x/crypto/bcrypt"
)

type jwtToken struct {
	Token string `json:"token"`
}

type priveteClaims struct {
	UserID    int    `json:"calico/user-id"`
	Mail      string `json:"calico/user-mail"`
	Name      string `json:"calico/user-name"`
	CreatedAt string `json:"calico/user-created_at"`
}

type myClaims struct {
	priveteClaims
	jwt.StandardClaims
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &myClaims{
		priveteClaims{
			// admmin?
			UserID:    bufDB.ID,
			Mail:      bufDB.Mail,
			Name:      bufDB.Name,
			CreatedAt: bufDB.CreatedAt,
		},
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

// OnlyPersonMiddleware ...
func OnlyPersonMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonString, err := jwtmiddleware.FromAuthHeader(r)
		if err != nil {
			log.Fatal(err)
		}

		var claims myClaims
		_, err = jwt.ParseWithClaims(jsonString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SIGNINKEY")), nil
		})
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

// AuthTest ...
var AuthTest = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	headerAuthorization := strings.Split(r.Header.Get("Authorization"), " ")

	// check "Bearer"

	var claims myClaims
	token, _ := jwt.ParseWithClaims(headerAuthorization[1], &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINKEY")), nil
	})
	fmt.Println(token.Valid)

	json.NewEncoder(w).Encode(claims)
})
