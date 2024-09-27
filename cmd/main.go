package main

import (
	"api/internal/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Iniciando Api")

	r := router.Router()

	log.Fatal(http.ListenAndServe(":4000", r))
}
