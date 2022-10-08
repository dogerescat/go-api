package routes

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func GetTodoRouter(todos *gin.RouterGroup) {
	todoController := controllers.NewTodoController()
	todos.POST("/", todoController.Create)
	todos.GET("/:id", todoController.GetAll)
}
