package controllers

import (
	"go-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoController struct{}

func NewTodoController() *TodoController {
	return &TodoController{}
}

func (tc *TodoController) Create(c *gin.Context) {
	title := c.PostForm("title")
	user_id, err := strconv.Atoi(c.PostForm("user_id"))
	if err != nil {
		panic(err)
	}
	data := models.Todo{Title: title, IsDone: false, UserId: user_id}
	data.Create()
	c.JSON(http.StatusOK, gin.H{
		"message": "create todo",
	})
}

func (tc *TodoController) GetAll(c *gin.Context) {
	id := c.Params.ByName("id")
	user_id, _ := strconv.Atoi(id)
	todos, err := models.GetAll(user_id)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, todos)
}
