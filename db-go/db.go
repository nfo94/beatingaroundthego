package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// local mysql server container exposed in port 3306
	stringConnection := "root:12345@tcp(0.0.0.0:3306)/mysqlgo?charset=utf8&parseTime=True&loc=Local"
	// to open the connection and create a var for errors
	db, err := sql.Open("mysql", stringConnection)
	// check error
	if err != nil {
		log.Fatal(err)
	}

	// ensure connections are closed
	db.SetConnMaxLifetime(time.Minute * 3)
	// limit the number of connections
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// to test connection
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Open database connection")

	lines, err := db.Query("create table users (id int, name varchar(50))")
	if err != nil {
		log.Fatal(err)
	}

	defer lines.Close()

	fmt.Println(lines)
}
