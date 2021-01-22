package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // driver to connect with MySQL
)

// Connect with database
func Connect() (*sql.DB, error) {
	stringConnection := "root:12345@tcp(0.0.0.0:3306)/gocrud?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to gocrud database")

	return db, nil
}
