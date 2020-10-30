package main

import (
	"log"

	"github.com/N4n3x/CI-Cloud-GO/config"
	"github.com/N4n3x/CI-Cloud-GO/routes"
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
