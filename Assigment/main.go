package main

import (
	"Assigment/routes"
	"fmt"
	"net/http"
)

// @title Golang API endpoints assignment
// @version 1.0
// @description This is a sample of Golang API endpoints assignment
// @termsOfService http://swagger.io/terms/

// @contact.name Hao Nguyen
// @contact.url http://musicmaven.s3corp.vn
// @contact.email hao.nguyen@s3corp.com.vn

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	fmt.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", routes.CreateRouter())
}
