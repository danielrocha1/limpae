package routes

import (
	"limpae/go/src/handlers"
	"github.com/gofiber/fiber/v2"
)

func ControllersRoutes(app *fiber.App) {
	api := app.Group("/users")

	api.Post("/", handlers.CreateUser)
	api.Get("/", handlers.GetUsers)
	api.Get("/:id", handlers.GetUser)
	api.Put("/:id", handlers.UpdateUser)
	api.Delete("/:id", handlers.DeleteUser)
}
