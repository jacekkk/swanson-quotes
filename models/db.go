package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(user string, password string, socket string, connectionName string, dbName string) {

	// /cloudsql is used on App Engine.
	if socket == "" {
		socket = "/cloudsql"
	}

	// MySQL Connection, comment out to use PostgreSQL.
	dbURI := fmt.Sprintf("%s:%s@unix(%s/%s)/%s", user, password, socket, connectionName, dbName)

	var err error
	db, err = sql.Open("mysql", dbURI)

	if err != nil {
		panic(fmt.Sprintf("DB: %v", err))
	}

	fmt.Println("db connected")
}
