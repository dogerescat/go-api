package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	users := r.Group("/users")
	{
		users.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello world",
			})
		})
		users.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "login",
			})
		})
	}
	r.Run(":80")
}
