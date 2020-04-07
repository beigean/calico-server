package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type todoCols struct {
	ID     int    `db:"id" json:"id"`
	UserID int    `db:"user_id" json:"user_id"`
	Todo   string `db:"todo" json:"todo"`
}

type todolist []todoCols

// TodoGet : get todos data list
func TodoGet(c *gin.Context) {
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

	c.JSON(http.StatusOK, buflist)
	return
}

// TodoGetByID : get todos data by id
func TodoGetByID(c *gin.Context) {
	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	var buf todoCols
	err = db.QueryRowx("SELECT * FROM todos WHERE id=?", c.Param("id")).StructScan(&buf)
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

// TodoPost : post new data to todos table
func TodoPost(c *gin.Context) {
	var buf todoCols
	err := c.BindJSON(&buf)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec("INSERT INTO todos (user_id, todo) VALUES (?,?)", buf.UserID, buf.Todo)
	fmt.Println(res)
	if err != nil {
		log.Fatal(err)
	}

	// c.JSON(http.StatusOK, nil)
	return
}

// TodoPut : update todo data by id
func TodoPut(c *gin.Context) {
	var buf todoCols
	err := c.BindJSON(&buf)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec("UPDATE todos SET user_id=?, todo=? WHERE id=?", buf.UserID, buf.Todo, c.Param("id"))
	fmt.Println(res)
	if err != nil {
		log.Fatal(err)
	}

	// c.JSON(http.StatusOK, nil)
	return
}

// TodoDelete : delete todo data by id
func TodoDelete(c *gin.Context) {
	db, err := sqlx.Open(kindDb, dsn)
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec("DELETE FROM todos WHERE id=?", c.Param("id"))
	fmt.Println(res)
	if err != nil {
		log.Fatal(err)
	}

	// c.JSON(http.StatusOK, nil)
	return
}
