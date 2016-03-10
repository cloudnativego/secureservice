package main

import (
	"log"
	"os"

	"github.com/cloudnativego/secureservice/server"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3001"
	}

	//short-circuit fail if api key is not set
	apikey := os.Getenv(server.APIKey)
	if len(apikey) == 0 {
		log.Fatal("Application is not properly configured.")
		os.Exit(1)
	}

	s := server.NewServer()
	s.Run(":" + port)
}
