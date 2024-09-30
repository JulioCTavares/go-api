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

	_, err := db.ConnectDB()

	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	err = db.RunMigrations()

	if err != nil {
		log.Fatal("Falha ao rodar a migração:", err)
	}

	log.Fatal(http.ListenAndServe(":4000", r))
}
