package handlers

import (
	"limpae/go/src/config"
	"limpae/go/src/models"
	"github.com/gofiber/fiber/v2"
)

// Criar Serviço
func CreateService(c *fiber.Ctx) error {
	service := new(models.Service)
	if err := c.BodyParser(service); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Create(&service)
	return c.JSON(service)
}

// Listar Serviços
func GetServices(c *fiber.Ctx) error {
	var services []models.Service
	config.DB.Find(&services)
	return c.JSON(services)
}

// Buscar Serviço por ID
func GetService(c *fiber.Ctx) error {
	id := c.Params("id")
	var service models.Service
	if err := config.DB.First(&service, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Serviço não encontrado"})
	}
	return c.JSON(service)
}

// Atualizar Serviço
func UpdateService(c *fiber.Ctx) error {
	id := c.Params("id")
	var service models.Service
	if err := config.DB.First(&service, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Serviço não encontrado"})
	}
	if err := c.BodyParser(&service); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Save(&service)
	return c.JSON(service)
}

// Deletar Serviço
func DeleteService(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.Service{}, id)
	return c.SendStatus(204)
}
