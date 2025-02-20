package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"limpae/go/src/config"
	"limpae/go/src/models"
	"limpae/go/src/utils"
)

// UploadPhotoHandler lida com o upload de fotos
func UploadPhotoHandler(c *fiber.Ctx) error {
	// Pegar o ID do usuário a partir dos parâmetros ou autenticação (ajuste conforme necessário)
	userIDParam := c.Params("userID") // Exemplo: pega da URL "/upload/:userID"
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID de usuário inválido"})
	}

	// Obter o arquivo do request
	file, err := c.FormFile("photo")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Arquivo inválido"})
	}

	// Abrir o arquivo
	fileContent, err := file.Open()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao abrir o arquivo"})
	}
	defer fileContent.Close()

	// Obter a extensão do arquivo (exemplo: .jpg, .png)
	fileExt := ""
	if len(file.Filename) > 4 {
		fileExt = file.Filename[len(file.Filename)-4:]
	}

	// Gerar nome do arquivo: "user-{ID}.ext"
	fileName := fmt.Sprintf("user-%d%s", userID, fileExt)

	// Fazer upload para o Supabase
	fileURL, err := utils.UploadFileToSupabase(fileContent, fileName)
	if err != nil {
		log.Println("Erro ao enviar para o Supabase:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao fazer upload"})
	}

	// Atualizar o campo Photo no banco de dados
	result := config.DB.Model(&models.User{}).Where("id = ?", userID).Update("photo", fileURL)
	if result.Error != nil {
		log.Println("Erro ao atualizar URL no banco:", result.Error)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao salvar URL da foto"})
	}

	// Retornar a URL da imagem
	return c.JSON(fiber.Map{"url": fileURL})
}
