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

	http.Handle("/home", middleware.AuthMiddleware(enableCors(http.HandlerFunc(handlers.Hello))))
	http.Handle("/register", enableCors(http.HandlerFunc(controllers.RegisterHandler)))
	http.Handle("/login", enableCors(http.HandlerFunc(controllers.LoginHandler)))
	http.Handle("/logout", middleware.AuthMiddleware(enableCors(http.HandlerFunc(controllers.LogoutHandler))))
	http.ListenAndServe(":8000", nil)
}

func enableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		h.ServeHTTP(w, r)
	})
}
