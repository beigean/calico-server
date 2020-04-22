package main

import (
	"log"
	"net/http"

	"github.com/noobs9/calico-server/pkg/auth"
	"github.com/noobs9/calico-server/pkg/controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/user/{id}", controller.UserGetByID).Methods("GET")
	r.Handle("/user", controller.UserGet).Methods("GET")
	r.Handle("/user", controller.UserPost).Methods("POST")
	r.Handle("/user/{id}", controller.UserPut).Methods("PUT")
	r.Handle("/user/{id}", controller.UserDelete).Methods("DELETE")
	r.Handle("/todo/{id}", controller.TodoGetByID).Methods("GET")
	r.Handle("/todo", controller.TodoGet).Methods("GET")
	r.Handle("/todo", controller.TodoPost).Methods("POST")
	r.Handle("/todo/{id}", controller.TodoPut).Methods("PUT")
	r.Handle("/todo/{id}", controller.TodoDelete).Methods("DELETE")

	r.HandleFunc("/auth", auth.GetTokenHandler).Methods("POST")
	r.Handle("/auth/test", auth.JwtMiddleware.Handler(auth.AuthTest))

	r.HandleFunc("/ping", pingHandler)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe failed: ", err)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\n"))
}
