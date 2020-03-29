package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Age  int    `db:"age" json:"age"`
}

type Userlist []User

func main() {
	r := gin.Default()
	r.GET("/ping", sample)
	r.GET("/todo/:id", todo_get)
	// r.GET("/todo", todo_get)
	r.POST("/todo", todo_post)
	r.PUT("/todo/:id", todo_put)
	r.DELETE("/todo/:id", todo_delete)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func sample(c *gin.Context) {
	// log.Printf()
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func todo_get(c *gin.Context) {
	var userlist Userlist

	// dsn spec: "[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]"
	db, err := sqlx.Open("mysql", "root:asn10026900@/calico")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Queryx("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	var user User
	for rows.Next() {

		err := rows.StructScan(&user)
		if err != nil {
			log.Fatal(err)
		}
		userlist = append(userlist, user)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(200, userlist[id])
}

func todo_post(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "post dayo",
	})
}

func todo_put(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "put dayo",
	})
}

func todo_delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete dayo",
	})
}
