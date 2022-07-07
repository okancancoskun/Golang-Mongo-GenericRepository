package main

import (
	"myapp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	routes.ProductRoutes(router)
	routes.UserRoutes(router)
	routes.AuthRoutes(router)
	router.Run()
}
