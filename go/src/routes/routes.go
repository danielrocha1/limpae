package routes

import (
	"limpae/go/src/handlers"
	"limpae/go/src/controllers"
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
	api.Post("/diarists", handlers.CreateDiaristProfile)
	// api.Get("/diarists", handlers.GetDiarists)
	// api.Get("/diarists/:id", handlers.GetDiarist)
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

	// Rotas para Assinaturas
	api.Post("/subscriptions", handlers.CreateSubscription)
	api.Get("/subscriptions", handlers.GetSubscriptions)
	api.Get("/subscriptions/:id", handlers.GetSubscription)
	api.Put("/subscriptions/:id", handlers.UpdateSubscription)
	api.Delete("/subscriptions/:id", handlers.CancelSubscription)

	api.Post("/userprofile", handlers.CreateUserProfile)
	api.Post("/upload-photo", controllers.UploadPhotoHandler)


	api.Get("/diarists-nearby", func(c *fiber.Ctx) error {
		return handlers.GetNearbyDiarists(c)
	})

	// Rota de fallback
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("Página não encontrada")
	})


}
