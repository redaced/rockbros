package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/redaced/rockbros/pkg/database"
	"github.com/redaced/rockbros/pkg/models"

	"gorm.io/gorm"
)

// Define database client
var db *gorm.DB = database.ConnectDB()

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := models.RegisterUser(db, creds.Username, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// At this point, the user is registered. You can set a cookie or a token for the user.
	json.NewEncoder(w).Encode(user)
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := models.AuthenticateUser(db, creds.Username, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	tokenString, err := models.GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(tokenString))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Get the session from the request
	// session, err := store.Get(r, "session-name")
	// if err != nil {
	//     http.Error(w, err.Error(), http.StatusInternalServerError)
	//     return
	// }

	// // Delete the session
	// session.Options.MaxAge = -1

	// err = session.Save(r, w)
	// if err != nil {
	//     http.Error(w, err.Error(), http.StatusInternalServerError)
	//     return
	// }

	// // Redirect the user to the login page (or wherever you want)
	// http.Redirect(w, r, "/login", http.StatusSeeOther)
}
