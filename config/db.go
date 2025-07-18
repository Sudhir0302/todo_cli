package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Conn() {
	var err error
	DB, err = sql.Open("sqlite", "./todo.db")
	if err != nil {
		log.Fatal(err)
	}
	if DB != nil {
		fmt.Println("db connected")
	}
}
