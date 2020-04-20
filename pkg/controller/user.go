package controller

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type userCols struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Age  int    `db:"age" json:"age"`
}

type userlist []userCols

// UserGet : get users data list
var UserGet = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Queryx("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var buf userCols
	var buflist userlist
	for rows.Next() {

		err := rows.StructScan(&buf)
		if err != nil {
			log.Fatal(err)
		}
		buflist = append(buflist, buf)
	}

	json.NewEncoder(w).Encode(buflist)
})

// UserGetByID : get users data by id
var UserGetByID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	var buf userCols
	err = db.QueryRowx("SELECT * FROM users WHERE id=?", mux.Vars(r)["id"]).StructScan(&buf)
	if err == nil {
		// pass
	} else if err == sql.ErrNoRows {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(buf)
})

// UserPost : post new data to users table
var UserPost = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	var buf userCols
	err = json.Unmarshal(body, &buf)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO users (name, age) VALUES (?,?)", buf.Name, buf.Age)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	return
})

// UserPut : update user data by id
var UserPut = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	var user userCols
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("UPDATE users SET name=?, age=? WHERE id=?", user.Name, user.Age, mux.Vars(r)["id"])
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
})

// UserDelete : delete user data by id
var UserDelete = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM users WHERE id=?", mux.Vars(r)["id"])
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
})
