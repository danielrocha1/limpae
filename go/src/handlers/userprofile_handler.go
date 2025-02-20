package handlers

import (
	"limpae/go/src/config"
	"limpae/go/src/models"
	"github.com/gofiber/fiber/v2"
)

// Criar User Profile
func CreateUserProfile(c *fiber.Ctx) error {
	userprofile := new(models.UserProfile)
	if err := c.BodyParser(userprofile); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := config.DB.Create(userprofile).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create userprofile"})
	}
	return c.Status(201).JSON(userprofile)
}

// Listar Todos os userprofileas
func GetUserProfiles(c *fiber.Ctx) error {
	var userprofiles []models.UserProfile
	config.DB.Find(&userprofiles)
	return c.JSON(userprofiles)
}

// Obter userprofilea por ID
func GetUserProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	var userprofile models.UserProfile
	if err := config.DB.First(&userprofile, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "userprofile not found"})
	}
	return c.JSON(userprofile)
}

// Atualizar userprofilea
func UpdateUserProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	var userprofile models.UserProfile
	if err := config.DB.First(&userprofile, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	if err := c.BodyParser(&userprofile); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	config.DB.Save(&userprofile)
	return c.JSON(userprofile)
}

// Deletar 
func DeleteUserProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.UserProfile{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete userprofile"})
	}
	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}
