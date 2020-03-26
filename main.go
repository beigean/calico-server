package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", sample)
	r.GET("/todo", todo_get)
	r.POST("/todo", todo_post)
	r.PUT("/todo", todo_put)
	r.DELETE("/todo", todo_delete)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func sample(c *gin.Context) {
	// log.Printf()
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func todo_get(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get dayo",
	})
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
