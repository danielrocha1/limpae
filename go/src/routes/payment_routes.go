package routes

import (
    "github.com/gofiber/fiber/v2"
    "limpae/go/src/handlers"
)

func SetupPaymentRoutes(api fiber.Router) {
    payments := api.Group("/payments")

    payments.Post("/", handlers.CreatePayment)
    payments.Get("/", handlers.GetPayments)
    payments.Get("/:id", handlers.GetPayment)
    payments.Put("/:id", handlers.UpdatePayment)
    payments.Delete("/:id", handlers.DeletePayment)
}
