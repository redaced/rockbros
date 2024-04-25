package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Username string
	Password string
}

var jwtKey = []byte("your_secret_key")

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func ParseJWT(tokenString string) (string, error) {
	claims := &jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
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
func ConnectDB(dataSourceName string) error {
	var err error
	DB, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	DB.AutoMigrate(&User{})
	return nil
}
func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func RegisterUser(username, password string) (*User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := User{Username: username, Password: string(hashedPassword)}
	if err := DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func AuthenticateUser(username, password string) (*User, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	return user, nil
}
