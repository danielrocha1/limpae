package main

import (
	"limpae/go/src/config"
	"limpae/go/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	// Iniciar conexão com o banco de dados
	config.ConnectDB()

	// Criar nova instância do Fiber
	app := fiber.New()

	// Middlewares globais
	app.Use(logger.New()) // Logger de requisições
	app.Use(cors.New())   // CORS para permitir requisições externas

	// Configurar Rotas
	routes.SetupRoutes(app)

	// Rota de teste
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 API de Diaristas rodando!")
	})

	// Iniciar servidor
	port := ":8080"
	log.Println("🚀 Servidor rodando na porta", port)
	if err := app.Listen(port); err != nil {
		log.Fatal(err)
	}
}
