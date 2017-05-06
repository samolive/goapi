package main

import (
	"goapi/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.Index)
	router.HandleFunc("/todos", controllers.TodoIndex)
	router.HandleFunc("/todos/{todoId}", controllers.TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}
