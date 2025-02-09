package handlers

import (
	"limpae/go/src/config"
	"limpae/go/src/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Criar nova assinatura
func CreateSubscription(c *fiber.Ctx) error {
	sub := new(models.Subscription)

	if err := c.BodyParser(sub); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Dados inválidos"})
	}

	// Definir preço do plano
	switch sub.Plan {
	case "free":
		sub.Price = 0.00
	case "basic":
		sub.Price = 29.99
	case "premium":
		sub.Price = 59.99
	default:
		return c.Status(400).JSON(fiber.Map{"error": "Plano inválido"})
	}

	sub.Status = "active"
	sub.ExpiresAt = time.Now().AddDate(0, 1, 0) // Expira em 1 mês

	config.DB.Create(&sub)

	return c.Status(201).JSON(sub)
}

// Obter todas as assinaturas
func GetSubscriptions(c *fiber.Ctx) error {
	var subs []models.Subscription
	config.DB.Find(&subs)
	return c.JSON(subs)
}

// Obter assinatura por ID
func GetSubscription(c *fiber.Ctx) error {
	id := c.Params("id")
	var sub models.Subscription

	if result := config.DB.First(&sub, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Assinatura não encontrada"})
	}

	return c.JSON(sub)
}

// Atualizar assinatura
func UpdateSubscription(c *fiber.Ctx) error {
	id := c.Params("id")
	var sub models.Subscription

	if result := config.DB.First(&sub, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Assinatura não encontrada"})
	}

	if err := c.BodyParser(&sub); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Dados inválidos"})
	}

	config.DB.Save(&sub)
	return c.JSON(sub)
}

// Cancelar assinatura
func CancelSubscription(c *fiber.Ctx) error {
	id := c.Params("id")
	var sub models.Subscription

	if result := config.DB.First(&sub, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Assinatura não encontrada"})
	}

	sub.Status = "canceled"
	config.DB.Save(&sub)

	return c.JSON(fiber.Map{"message": "Assinatura cancelada com sucesso"})
}
