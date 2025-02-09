package routes

import (
    "github.com/gofiber/fiber/v2"
    "limpae/go/src/handlers"
)

func SetupReviewRoutes(api fiber.Router) {
    reviews := api.Group("/reviews")

    reviews.Post("/", handlers.CreateReview)
    reviews.Get("/", handlers.GetReviews)
    reviews.Get("/:id", handlers.GetReview)
    reviews.Put("/:id", handlers.UpdateReview)
    reviews.Delete("/:id", handlers.DeleteReview)
}
