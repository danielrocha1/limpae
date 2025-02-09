package handlers

import (
	"limpae/go/src/config"
	"limpae/go/src/models"
	"github.com/gofiber/fiber/v2"
)

// Criar Diarista
func CreateDiarist(c *fiber.Ctx) error {
	diarist := new(models.Diarist)
	if err := c.BodyParser(diarist); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := config.DB.Create(diarist).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create diarist"})
	}
	return c.Status(201).JSON(diarist)
}

// Listar Todos os Diaristas
func GetDiarists(c *fiber.Ctx) error {
	var diarists []models.Diarist
	config.DB.Find(&diarists)
	return c.JSON(diarists)
}

// Obter Diarista por ID
func GetDiarist(c *fiber.Ctx) error {
	id := c.Params("id")
	var diarist models.Diarist
	if err := config.DB.First(&diarist, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Diarist not found"})
	}
	return c.JSON(diarist)
}

// Atualizar Diarista
func UpdateDiarist(c *fiber.Ctx) error {
	id := c.Params("id")
	var diarist models.Diarist
	if err := config.DB.First(&diarist, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Diarist not found"})
	}
	if err := c.BodyParser(&diarist); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	config.DB.Save(&diarist)
	return c.JSON(diarist)
}

// Deletar Diarista
func DeleteDiarist(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Diarist{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete diarist"})
	}
	return c.JSON(fiber.Map{"message": "Diarist deleted successfully"})
}
