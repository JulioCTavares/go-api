package main

import (
	"api/internal/db"
	"api/internal/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Iniciando Api")

	r := router.Router()

	db.ConnectDB()

	log.Fatal(http.ListenAndServe(":4000", r))
}
