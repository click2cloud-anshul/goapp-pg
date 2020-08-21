package main

import (
	"fmt"
	"log"
	"net/http"

	"go-postgres1.1/router"
)

func main() {

	r := router.Router()
	fmt.Println("connecting to the port 8080...")

	// connection to server
	log.Fatal(http.ListenAndServe(":8080", r))

}
