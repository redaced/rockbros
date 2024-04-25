package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	// Replace the placeholders with your MySQL connection details
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/rockbros")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Connected to MySQL database!")

	return db, nil
}
