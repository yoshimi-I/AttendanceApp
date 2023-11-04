// main.go
package main

import (
	"net/http"
	"work-management-app/presentation/router"
)

func main() {
	r := router.Router()
	http.ListenAndServe(":8080", r)
}
