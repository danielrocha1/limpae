package handlers

import (
	"limpae/go/src/config"
	"limpae/go/src/models"

	"github.com/gofiber/fiber/v2"
)

// Criar Pagamento
func CreatePayment(c *fiber.Ctx) error {
	payment := new(models.Payment)
	if err := c.BodyParser(payment); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Create(&payment)
	return c.JSON(payment)
}

// Listar Pagamentos
func GetPayments(c *fiber.Ctx) error {
	var payments []models.Payment
	config.DB.Find(&payments)
	return c.JSON(payments)
}

// Buscar Pagamento por ID
func GetPayment(c *fiber.Ctx) error {
	id := c.Params("id")
	var payment models.Payment
	if err := config.DB.First(&payment, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Pagamento não encontrado"})
	}
	return c.JSON(payment)
}

// Atualizar Pagamento
func UpdatePayment(c *fiber.Ctx) error {
	id := c.Params("id")
	var payment models.Payment
	if err := config.DB.First(&payment, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Pagamento não encontrado"})
	}
	if err := c.BodyParser(&payment); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Save(&payment)
	return c.JSON(payment)
}

// Deletar Pagamento
func DeletePayment(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.Payment{}, id)
	return c.SendStatus(204)
}
