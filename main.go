package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"swanson/models"
	"swanson/router"
)

func main() {
	var (
		connectionName = mustGetenv("CLOUDSQL_CONNECTION_NAME")
		user           = mustGetenv("CLOUDSQL_USER")
		dbName         = os.Getenv("CLOUDSQL_DATABASE_NAME") // NOTE: dbName may be empty
		password       = os.Getenv("CLOUDSQL_PASSWORD")      // NOTE: password may be empty
		socket         = os.Getenv("CLOUDSQL_SOCKET_PREFIX")
		port           = os.Getenv("PORT")
	)

	models.InitDB(user, password, socket, connectionName, dbName)
	router := router.Router()

	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}
