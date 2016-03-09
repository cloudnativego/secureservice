package main

import (
	"os"

	"github.com/cloudnativego/secureservice/server"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3001"
	}

	s := server.NewServer()
	s.Run(":" + port)
}
