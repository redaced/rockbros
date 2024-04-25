package main

import (
	"log"
	"net/http"

	"github.com/redaced/rockbros/pkg/api/controllers"
	"github.com/redaced/rockbros/pkg/api/handlers"
	"github.com/redaced/rockbros/pkg/api/middleware"
	"github.com/redaced/rockbros/pkg/models"
)

func main() {
	err := models.ConnectDB("root:password@/rockbros?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/home", middleware.AuthMiddleware(http.HandlerFunc(handlers.Hello)))
	http.HandleFunc("/register", controllers.RegisterHandler)
	http.HandleFunc("/login", controllers.LoginHandler)
	http.ListenAndServe(":8080", nil)
}
