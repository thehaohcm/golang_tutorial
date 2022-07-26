package main

import (
	"golang_project/routes"
)

func main() {
	server := routes.SetupRouter()
	server.Run(":8080")
}
