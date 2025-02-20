package handlers

import (
	"limpae/go/src/config"
	"limpae/go/src/models"
	"net/http"
	"strconv"
	"fmt"


	"github.com/gofiber/fiber/v2"
	"github.com/umahmood/haversine"
)

// GetNearbyDiarists retorna diaristas próximos a uma localização específica
func GetNearbyDiarists(c *fiber.Ctx) error {
	// Captura os parâmetros da query
	latParam := c.Query("latitude")
	lonParam := c.Query("longitude")

	// Converte para float64
	latitude, err1 := strconv.ParseFloat(latParam, 64)
	longitude, err2 := strconv.ParseFloat(lonParam, 64)
	if err1 != nil || err2 != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Parâmetros de latitude e longitude inválidos"})
	}

	// Busca os diaristas e seus endereços
	var diarists []models.User
	err := config.DB.Preload("Address").Where("role = ?", "diarista").Find(&diarists).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao buscar diaristas"})
	}

	// Estrutura para armazenar diaristas com distância calculada
	type DiaristWithDistance struct {
		ID       uint    `json:"id"`
		Name     string  `json:"name"`
		Distance string `json:"distance"`
	}

	var nearbyDiarists []DiaristWithDistance

	// Calcula a distância usando Haversine
	origin := haversine.Coord{Lat: latitude, Lon: longitude}
	for _, diarist := range diarists {
		if len(diarist.Address) > 0 {
			addr := diarist.Address[0] // Pegamos o primeiro endereço cadastrado
			dest := haversine.Coord{Lat: addr.Latitude, Lon: addr.Longitude}
			_, km := haversine.Distance(origin, dest)
			distance := fmt.Sprintf("%.2f km", km) // Formata com duas casas decimais e adiciona "km"


			nearbyDiarists = append(nearbyDiarists, DiaristWithDistance{
				ID:       diarist.ID,
				Name:     diarist.Name,
				Distance: distance,
			})
		}
	}

	// Retorna a lista ordenada por distância
	return c.JSON(nearbyDiarists)
}
