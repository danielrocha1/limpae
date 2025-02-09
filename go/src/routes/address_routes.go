package routes

import (
    "github.com/gofiber/fiber/v2"
    "limpae/go/src/handlers"
)

func SetupAddressRoutes(api fiber.Router) {
    address := api.Group("/addresses")

    address.Post("/", handlers.CreateAddress)
    address.Get("/", handlers.GetAddresses)
    address.Get("/:id", handlers.GetAddress)
    address.Put("/:id", handlers.UpdateAddress)
    address.Delete("/:id", handlers.DeleteAddress)
}
