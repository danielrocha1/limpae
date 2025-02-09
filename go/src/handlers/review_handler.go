package handlers

import (
	"limpae/go/src/config"
	"limpae/go/src/models"
	"github.com/gofiber/fiber/v2"
)

// Criar Avaliação
func CreateReview(c *fiber.Ctx) error {
	review := new(models.Review)
	if err := c.BodyParser(review); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := config.DB.Create(review).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create review"})
	}
	return c.Status(201).JSON(review)
}

// Listar Todas as Avaliações
func GetReviews(c *fiber.Ctx) error {
	var reviews []models.Review
	config.DB.Find(&reviews)
	return c.JSON(reviews)
}

// Obter Avaliação por ID
func GetReview(c *fiber.Ctx) error {
	id := c.Params("id")
	var review models.Review
	if err := config.DB.First(&review, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Review not found"})
	}
	return c.JSON(review)
}

// Atualizar Avaliação
func UpdateReview(c *fiber.Ctx) error {
	id := c.Params("id")
	var review models.Review
	if err := config.DB.First(&review, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Review not found"})
	}
	if err := c.BodyParser(&review); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	config.DB.Save(&review)
	return c.JSON(review)
}

// Deletar Avaliação
func DeleteReview(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Review{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete review"})
	}
	return c.JSON(fiber.Map{"message": "Review deleted successfully"})
}
