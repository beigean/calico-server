package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type userCols struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Age  int    `db:"age" json:"age"`
}

type userlist []userCols

const (
	kindDb = "mysql"
	// dsn spec: "[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]"
	dsn = "root:asn10026900@/calico"
)

// UserGet : get users data list
func UserGet(c *gin.Context) {
	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Queryx("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	var user userCols
	var userlist userlist
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

// UserGetByID : get users data by id
func UserGetByID(c *gin.Context) {
	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Queryx("SELECT * FROM users WHERE id=?", c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	var user userCols
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatal(err)
		}
	}

	c.JSON(http.StatusOK, user)
	return
}

// UserPost : post new data to users table
func UserPost(c *gin.Context) {
	var user userCols
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

// UserPut : update user data by id
func UserPut(c *gin.Context) {
	var user userCols
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec("UPDATE users SET name=?, age=? WHERE id=?", user.Name, user.Age, c.Param("id"))
	fmt.Println(res)
	if err != nil {
		log.Fatal(err)
	}

	// c.JSON(http.StatusOK, nil)
	return
}

// UserDelete : delete user data by id
func UserDelete(c *gin.Context) {
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
