package server

import (
	"fmt"
	"go-api/routes"

	"github.com/gin-gonic/gin"
)

func Run(port int) {
	run(port)
}

func run(port int) {
	r := gin.Default()
	users := r.Group("/users")
	routes.GetRouter(users)
	r.Run(fmt.Sprintf(":%d", port))
}
