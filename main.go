package main

import (
	"CI-Cloud-GO/config"
	"CI-Cloud-GO/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Database
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
