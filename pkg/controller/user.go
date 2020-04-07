package controller

import (
	"database/sql"
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

	c.JSON(http.StatusOK, buflist)
	return
}

// UserGetByID : get users data by id
func UserGetByID(c *gin.Context) {
	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	var buf userCols
	err = db.QueryRowx("SELECT * FROM users WHERE id=?", c.Param("id")).StructScan(&buf)
	if err == nil {
		// pass
	} else if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id."})
		return
	} else {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, buf)
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
