package main

import (
	"net/http"

	handlers "github.com/redaced/rockbros/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Hello)
	http.ListenAndServe(":8080", nil)
}
