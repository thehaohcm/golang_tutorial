package main

import (
	"Assigment/routes"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", routes.CreateRouter())
}
