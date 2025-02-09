package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"limpae/go/src/models"
)

var DB *gorm.DB

func ConnectDB() {
	// Carregar variáveis de ambiente
	err := godotenv.Load("./src/config/.env")
	if err != nil {
		log.Fatal("Erro ao carregar o .env:", err)
	}

	// Pegar a variável DATABASE_URL
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("A variável de ambiente DATABASE_URL não está configurada.")
	}

	// Conectar ao banco de dados
	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	// Realizar AutoMigrate para criar/atualizar tabelas
	err = DB.AutoMigrate(
		&models.User{},
		&models.Address{},
		&models.Diarist{},
		&models.Service{},
		&models.Payment{},
		&models.Review{},
		&models.Subscription{}, // Adicionado
	)
	if err != nil {
		log.Fatal("Erro ao migrar tabelas:", err)
	}

	fmt.Println("✅ Banco de dados conectado e migrado com sucesso!")
}
