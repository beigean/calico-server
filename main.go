package main

import (
	"fmt"
	"log"
	"net/http"

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

const kindDb = "mysql"

// dsn spec: "[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]"
const dsn = "root:asn10026900@/calico"

func main() {
	r := gin.Default()
	r.GET("/todo/:id", todoGetByID)
	r.GET("/todo", todoGet)
	r.POST("/todo", todoPost)
	r.PUT("/todo/:id", todoPut)
	r.DELETE("/todo/:id", todoDelete)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func todoGet(c *gin.Context) {
	var userlist Userlist

	db, err := sqlx.Open(kindDb, dsn)
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

	c.JSON(http.StatusOK, userlist)
	return
}

func todoGetByID(c *gin.Context) {
	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Queryx("SELECT * FROM users WHERE id=?", c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	var user User
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatal(err)
		}
	}

	c.JSON(http.StatusOK, user)
	return
}

func todoPost(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec("INSERT INTO users (name, age) VALUES (?,?)", user.Name, user.Age)
	fmt.Println(res)
	if err != nil {
		log.Fatal(err)
	}

	// c.JSON(http.StatusOK, nil)
	return
}

func todoPut(c *gin.Context) {
	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec("UPDATE FROM users WHERE id=? LIMIT 1", c.Param("id"))
	fmt.Println(res)
	if err != nil {
		log.Fatal(err)
	}

	// c.JSON(http.StatusOK, nil)
	return
}

func todoDelete(c *gin.Context) {
	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec("DELETE FROM users WHERE id=?", c.Param("id"))
	fmt.Println(res)
	if err != nil {
		log.Fatal(err)
	}

	// c.JSON(http.StatusOK, nil)
	return
}
