package db

import (
	"fmt"
	"log"
	"os"

	"api/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading.env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		host, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("falha ao conectar ao banco de dados: %w", err)
	}

	fmt.Println("Conexão com o banco de dados bem-sucedida!")
	DB = db
	return db, nil
}

func RunMigrations() error {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		return fmt.Errorf("falha ao rodar a migração: %w", err)
	}
	fmt.Println("Migrações aplicadas com sucesso!")
	return nil
}
