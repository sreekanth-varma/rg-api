package main

import (
	"log"

	"github.com/gin-gonic/gin"
	server "github.com/sreekanth-varma/rg-core/rg-server"
)

func main() {

	app := gin.Default()
	options := server.GetDefaultOptions()
	options.WebServerPreHandler = configureRoutes
	server.Start(options)

	// Start the server
	log.Println("server is started on port 8080...")
	if err := app.Run(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}

func configureRoutes(router *gin.Engine) {

}
