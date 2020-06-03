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
		port           = os.Getenv("PORT")
		env            = mustGetenv("ENVIRONMENT")
	)

	models.InitDB(user, password, connectionName, dbName, env)
	router := router.Router()

	fmt.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}
