package main

import (
	"goapi/routes"
	"log"
	"net/http"
)

func main() {

	router := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
