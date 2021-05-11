package main

import (
	"golang-restful/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	routes.SetupRoutes(router)
}
