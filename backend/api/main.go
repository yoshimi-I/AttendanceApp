// main.go
package main

import (
	"github.com/yoshimi-I/AttendanceApp/presentation/router"
	"net/http"
)

func main() {
	r := router.Router()
	http.ListenAndServe(":8080", r)
}
