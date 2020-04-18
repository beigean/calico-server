package main

import (
	"log"
	"net/http"

	"github.com/noobs9/calico-server/pkg/controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", controller.UserGetByID).Methods("GET")
	r.HandleFunc("/user", controller.UserGet).Methods("GET")
	r.HandleFunc("/user", controller.UserPost).Methods("POST")
	r.HandleFunc("/user/{id}", controller.UserPut).Methods("PUT")
	r.HandleFunc("/user/{id}", controller.UserDelete).Methods("DELETE")
	r.HandleFunc("/todo/{id}", controller.TodoGetByID).Methods("GET")
	r.HandleFunc("/todo", controller.TodoGet).Methods("GET")
	r.HandleFunc("/todo", controller.TodoPost).Methods("POST")
	r.HandleFunc("/todo/{id}", controller.TodoPut).Methods("PUT")
	r.HandleFunc("/todo/{id}", controller.TodoDelete).Methods("DELETE")
	r.HandleFunc("/ping", pingHandler)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe failed: ", err)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\n"))
}
