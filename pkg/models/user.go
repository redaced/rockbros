package models

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	// Add more fields as needed
}

func NewUser(name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

// Add additional methods or functions related to the user model here
