package main

import (
	"go-app-postgres/cmd/config"
	"go-app-postgres/internal/router"
)

func main() {

	server := router.StartServer()

	server.Run(config.Port)

}
