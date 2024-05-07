package main

import (
	"gorm.io/gorm"

	"github.com/redaced/rockbros/pkg/database"
	"github.com/redaced/rockbros/pkg/routes"
)

var (
	db *gorm.DB = database.ConnectDB()
)

func main() {
	defer database.DisconnectDB(db)
	routes.SetupRoutes()
}
