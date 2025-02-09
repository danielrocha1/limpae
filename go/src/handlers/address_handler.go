package handlers

import (
	"limpae/go/src/config"
	"limpae/go/src/models"
	"github.com/gofiber/fiber/v2"
)

// Criar Endereço
func CreateAddress(c *fiber.Ctx) error {
	address := new(models.Address)
	if err := c.BodyParser(address); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Create(&address)
	return c.JSON(address)
}

// Listar Endereços
func GetAddresses(c *fiber.Ctx) error {
	var addresses []models.Address
	config.DB.Find(&addresses)
	return c.JSON(addresses)
}

// Buscar Endereço por ID
func GetAddress(c *fiber.Ctx) error {
	id := c.Params("id")
	var address models.Address
	if err := config.DB.First(&address, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Endereço não encontrado"})
	}
	return c.JSON(address)
}

// Atualizar Endereço
func UpdateAddress(c *fiber.Ctx) error {
	id := c.Params("id")
	var address models.Address
	if err := config.DB.First(&address, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Endereço não encontrado"})
	}
	if err := c.BodyParser(&address); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Save(&address)
	return c.JSON(address)
}

// Deletar Endereço
func DeleteAddress(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.Address{}, id)
	return c.SendStatus(204)
}
