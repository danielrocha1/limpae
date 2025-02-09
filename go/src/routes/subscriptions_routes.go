package routes

import (
	"limpae/go/src/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupSubscriptionRoutes(api fiber.Router) {
	subscriptionRoutes := api.Group("/subscriptions")

	subscriptionRoutes.Post("/", handlers.CreateSubscription)
	subscriptionRoutes.Get("/", handlers.GetSubscriptions)
	subscriptionRoutes.Get("/:id", handlers.GetSubscription)
	subscriptionRoutes.Put("/:id", handlers.UpdateSubscription)
	subscriptionRoutes.Delete("/:id", handlers.CancelSubscription)
}
