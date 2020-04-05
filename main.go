package main

import (
	"github.com/noobs9/calico-server/pkg/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	r.GET("/user/:id", controller.UserGetByID)
	r.GET("/user", controller.UserGet)
	r.POST("/user", controller.UserPost)
	r.PUT("/user/:id", controller.UserPut)
	r.DELETE("/user/:id", controller.UserDelete)
	r.GET("/todo/:id", controller.TodoGetByID)
	r.GET("/todo", controller.TodoGet)
	r.POST("/todo", controller.TodoPost)
	r.PUT("/todo/:id", controller.TodoPut)
	r.DELETE("/todo/:id", controller.TodoDelete)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
