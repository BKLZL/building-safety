package main

import (
	"building-safety/config"
	"building-safety/routes"
)

func main() {
	config.ConnectDatabase()

	r := routes.SetupRouter()

	r.Run(":8080")
}
