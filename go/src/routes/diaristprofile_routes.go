package routes

import (
    "github.com/gofiber/fiber/v2"
    "limpae/go/src/handlers"
)

func SetupDiaristRoutes(api fiber.Router) {
    diarists := api.Group("/diarists")

    diarists.Post("/", handlers.CreateDiaristProfile)
    diarists.Get("/", handlers.GetDiarists)
    diarists.Get("/:id", handlers.GetDiarist)
    diarists.Put("/:id", handlers.UpdateDiarist)
    diarists.Delete("/:id", handlers.DeleteDiarist)
}
