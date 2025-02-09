package routes

import (
	"limpae/go/src/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Rotas para Usuários
	api.Post("/users", handlers.CreateUser)
	api.Get("/users", handlers.GetUsers)
	api.Get("/users/:id", handlers.GetUser)
	api.Put("/users/:id", handlers.UpdateUser)
	api.Delete("/users/:id", handlers.DeleteUser)

	// Rotas para Diaristas
	api.Post("/diarists", handlers.CreateDiarist)
	api.Get("/diarists", handlers.GetDiarists)
	api.Get("/diarists/:id", handlers.GetDiarist)
	api.Put("/diarists/:id", handlers.UpdateDiarist)
	api.Delete("/diarists/:id", handlers.DeleteDiarist)

	// Rotas para Endereços
	api.Post("/addresses", handlers.CreateAddress)
	api.Get("/addresses", handlers.GetAddresses)
	api.Get("/addresses/:id", handlers.GetAddress)
	api.Put("/addresses/:id", handlers.UpdateAddress)
	api.Delete("/addresses/:id", handlers.DeleteAddress)

	// Rotas para Serviços
	api.Post("/services", handlers.CreateService)
	api.Get("/services", handlers.GetServices)
	api.Get("/services/:id", handlers.GetService)
	api.Put("/services/:id", handlers.UpdateService)
	api.Delete("/services/:id", handlers.DeleteService)

	// Rotas para Pagamentos
	api.Post("/payments", handlers.CreatePayment)
	api.Get("/payments", handlers.GetPayments)
	api.Get("/payments/:id", handlers.GetPayment)
	api.Put("/payments/:id", handlers.UpdatePayment)
	api.Delete("/payments/:id", handlers.DeletePayment)

	// Rotas para Avaliações
	api.Post("/reviews", handlers.CreateReview)
	api.Get("/reviews", handlers.GetReviews)
	api.Get("/reviews/:id", handlers.GetReview)
	api.Put("/reviews/:id", handlers.UpdateReview)
	api.Delete("/reviews/:id", handlers.DeleteReview)
}
