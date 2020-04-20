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

type todoCols struct {
	ID     int    `db:"id" json:"id"`
	UserID int    `db:"user_id" json:"user_id"`
	Todo   string `db:"todo" json:"todo"`
}

type todolist []todoCols

// TodoGet : get todos data list
var TodoGet = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Queryx("SELECT * FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var buf todoCols
	var buflist todolist
	for rows.Next() {
		err := rows.StructScan(&buf)
		if err != nil {
			log.Fatal(err)
		}
		buflist = append(buflist, buf)
	}

	json.NewEncoder(w).Encode(buflist)
})

// TodoGetByID : get todos data by id
var TodoGetByID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	var buf todoCols
	err = db.QueryRowx("SELECT * FROM todos WHERE id=?", mux.Vars(r)["id"]).StructScan(&buf)
	if err == nil {
		// pass
	} else if err == sql.ErrNoRows {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(buf)
})

// TodoPost : post new data to todos table
var TodoPost = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	var buf todoCols
	err = json.Unmarshal(body, &buf)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO todos (user_id, todo) VALUES (?,?)", buf.UserID, buf.Todo)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
})

// TodoPut : update todo data by id
var TodoPut = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	var buf todoCols
	err = json.Unmarshal(body, &buf)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("UPDATE todos SET user_id=?, todo=? WHERE id=?", buf.UserID, buf.Todo, mux.Vars(r)["id"])
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
})

// TodoDelete : delete todo data by id
var TodoDelete = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM todos WHERE id=?", mux.Vars(r)["id"])
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
})
