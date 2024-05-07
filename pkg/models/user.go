package models

import (
	"os"
	"time"

	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	expire, err := strconv.Atoi(os.Getenv("JWT_TOKEN_EXPIRE"))
	if err != nil {
		return "", err
	}
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(expire)).Unix()
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func ParseJWT(tokenString string) (string, error) {
	claims := &jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("JWT_TOKEN"), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		username := (*claims)["username"].(string)
		return username, nil
	}

	return "", err
}

func GetUserByUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func RegisterUser(db *gorm.DB, username, password string) (*User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := User{Username: username, Password: string(hashedPassword)}
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func AuthenticateUser(db *gorm.DB, username, password string) (*User, error) {

	user, err := GetUserByUsername(db, username)

	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	return user, nil
}
