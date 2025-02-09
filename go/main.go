package main

import (
	"limpae/go/src/config"
	"limpae/go/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
	"fmt"
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

	// Rota inicial
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	// Pega a variável de ambiente 'PORT' fornecida pelo Render
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000" // Porta padrão para desenvolvimento local, caso a variável não esteja configurada
	}

	// Inicia o servidor e faz o bind para o host 0.0.0.0 e a porta fornecida
	address := fmt.Sprintf("0.0.0.0:%s", port) // Explicitando que estamos ouvindo em 0.0.0.0
	err := app.Listen(address)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}

	fmt.Printf("Example app listening on port %s\n", port)
}
