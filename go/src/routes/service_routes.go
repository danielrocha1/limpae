package routes

import (
    "github.com/gofiber/fiber/v2"
    "limpae/go/src/handlers"
)

func SetupServiceRoutes(api fiber.Router) {
    services := api.Group("/services")

    services.Post("/", handlers.CreateService)
    services.Get("/", handlers.GetServices)
    services.Get("/:id", handlers.GetService)
    services.Put("/:id", handlers.UpdateService)
    services.Delete("/:id", handlers.DeleteService)
}
