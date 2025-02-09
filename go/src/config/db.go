package config

import (
	"fmt"
	"os"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load("/home/warlord/Projetos/limpae/go/src/config/.env")
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("A variável de ambiente DATABASE_URL não está configurada.")
	}

	// Usando a URL diretamente para conectar ao banco de dados
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	DB = db
	fmt.Println("Banco conectado!")
}
