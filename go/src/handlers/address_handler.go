package handlers

import (
	"encoding/json"
	"fmt"
	"limpae/go/src/config"
	"limpae/go/src/models"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Validação de CEP (Formato 00000-000)
func isValidZipcode(zipcode string) bool {
	re := regexp.MustCompile(`^\d{5}-\d{3}$`)
	return re.MatchString(zipcode)
}

// Validação de Estado (UF)
func isValidState(state string) bool {
	validStates := map[string]bool{
		"AC": true, "AL": true, "AP": true, "AM": true, "BA": true, "CE": true,
		"DF": true, "ES": true, "GO": true, "MA": true, "MT": true, "MS": true,
		"MG": true, "PA": true, "PB": true, "PR": true, "PE": true, "PI": true,
		"RJ": true, "RN": true, "RS": true, "RO": true, "RR": true, "SC": true,
		"SP": true, "SE": true, "TO": true,
	}
	return validStates[state]
}

// Função para buscar latitude e longitude com base no endereço
func getCoordinates(address string) (float64, float64, error) {
	baseURL := "https://nominatim.openstreetmap.org/search"
	query := url.Values{}
	query.Set("q", address)
	query.Set("format", "json")

	resp, err := http.Get(fmt.Sprintf("%s?%s", baseURL, query.Encode()))
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var results []struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return 0, 0, err
	}

	if len(results) == 0 {
		return 0, 0, fmt.Errorf("endereço não encontrado")
	}

	// Converter latitude e longitude para float64
	var lat, lon float64
	fmt.Sscanf(results[0].Lat, "%f", &lat)
	fmt.Sscanf(results[0].Lon, "%f", &lon)

	return lat, lon, nil
}

// Criar Endereço com validação
func CreateAddress(c *fiber.Ctx) error {
	address := new(models.Address)
	if err := c.BodyParser(address); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Dados inválidos"})
	}

	// Remover espaços extras
	address.Street = strings.TrimSpace(address.Street)
	address.Number = strings.TrimSpace(address.Number)
	address.Neighborhood = strings.TrimSpace(address.Neighborhood)
	address.City = strings.TrimSpace(address.City)
	address.State = strings.ToUpper(strings.TrimSpace(address.State))
	address.Zipcode = strings.TrimSpace(address.Zipcode)

	// Validações
	if address.UserID == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "ID do usuário é obrigatório"})
	}
	if address.Street == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Rua é obrigatória"})
	}
	if address.Neighborhood == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Bairro é obrigatório"})
	}
	if address.City == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Cidade é obrigatória"})
	}
	if !isValidState(address.State) {
		return c.Status(400).JSON(fiber.Map{"error": "Estado inválido"})
	}
	if !isValidZipcode(address.Zipcode) {
		return c.Status(400).JSON(fiber.Map{"error": "CEP inválido, formato esperado: 00000-000"})
	}

	// Montar endereço completo para busca da latitude e longitude
	fullAddress := fmt.Sprintf("%s, %s, %s, %s, %s, Brasil", address.Street, address.Number, address.Neighborhood, address.City, address.State)
	fmt.Println(fullAddress)
	// Buscar latitude e longitude
	latitude, longitude, err := getCoordinates(fullAddress)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Não foi possível obter coordenadas para este endereço"})
	}

	address.Latitude = latitude
	address.Longitude = longitude

	// Verificar se o usuário existe antes de cadastrar o endereço
	var user models.User
	if err := config.DB.First(&user, address.UserID).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Usuário não encontrado"})
	}

	// Criar endereço no banco
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
