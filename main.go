package main

import (
	"fmt"
	"log"
	"net/http"

	"swanson/models"
	"swanson/router"
)

func main() {
	models.InitDB()
	router := router.Router()

	fmt.Println("Starting server on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
