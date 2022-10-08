package controllers

import (
	"go-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) Create(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	data := models.User{Name: name, Email: email, Password: password}
	data.Create()
	c.JSON(http.StatusOK, gin.H{
		"message": "create user",
	})
}
