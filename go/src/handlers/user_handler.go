package handlers

import (
	"limpae/go/src/config"
	"limpae/go/src/models"
	"github.com/gofiber/fiber/v2"
)

// Criar Usuário
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Create(&user)
	return c.JSON(user)
}

// Listar Usuários
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	config.DB.Find(&users)
	return c.JSON(users)
}

// Buscar Usuário por ID
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Usuário não encontrado"})
	}
	return c.JSON(user)
}

// Atualizar Usuário
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Usuário não encontrado"})
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Save(&user)
	return c.JSON(user)
}

// Deletar Usuário
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.User{}, id)
	return c.SendStatus(204)
}
