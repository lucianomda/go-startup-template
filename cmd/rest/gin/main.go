package main

import (
	"fmt"

	"github.com/lucianomda/go-startup-template/internal/infrastructure/entrypoints/gin/handlers/ping"
	"github.com/lucianomda/go-startup-template/internal/infrastructure/entrypoints/gin/server"
)

func main() {
	pingHandler := ping.NewHandler()
	ginServer := server.New(pingHandler)
	err := ginServer.Run(":8080")
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
}
