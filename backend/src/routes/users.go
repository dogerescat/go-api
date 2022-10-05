package routes

import (
	"go-api/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRouter(users *gin.RouterGroup) {
	userController := controllers.NewUserController()
	users.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})
	users.POST("/", userController.Create)
}
