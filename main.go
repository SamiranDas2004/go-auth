// main.go
package main

import (
	"net/http"

	"github.com/SamiranDas2004/go-auth/dbconnect"
	"github.com/SamiranDas2004/go-auth/routes"
)

func main() {
	// Initialize a new Fiber app
	dbconnect.ConnectMongoDB()
	r := routes.Setup()
	http.ListenAndServe(":4000", r)
}
