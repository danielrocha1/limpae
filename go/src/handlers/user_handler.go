package handlers

import (
	"fmt"
	"limpae/go/src/config"
	"limpae/go/src/models"
	"regexp"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Função para validar CPF
func isValidCPF(cpf string) bool {
	cpf = strings.TrimSpace(cpf)
	cpf = strings.ReplaceAll(strings.ReplaceAll(cpf, ".", ""), "-", "") // Remove pontos e traço

	if len(cpf) != 11 || todosDigitosIguais(cpf) {
		return false
	}

	digito1 := calcularDigito(cpf[:9], 10)
	digito2 := calcularDigito(cpf[:10], 11)

	return cpf[9:] == fmt.Sprintf("%d%d", digito1, digito2)
}

// Verifica se todos os dígitos do CPF são iguais
func todosDigitosIguais(cpf string) bool {
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}
	return true
}

// Calcula os dígitos verificadores do CPF
func calcularDigito(cpfParcial string, pesoInicial int) int {
	soma := 0
	for i, peso := 0, pesoInicial; i < len(cpfParcial); i, peso = i+1, peso-1 {
		num, _ := strconv.Atoi(string(cpfParcial[i]))
		soma += num * peso
	}

	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto
}

// Função para validar e-mail
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// Função para validar telefone
func isValidPhone(phone string) bool {
	phone = strings.TrimSpace(phone)
	phone = strings.ReplaceAll(strings.ReplaceAll(phone, "(", ""), ")", "") // Remove parênteses
	phone = strings.ReplaceAll(phone, "-", "")                             // Remove hífen

	re := regexp.MustCompile(`^\d{10,11}$`) // Aceita números com 10 ou 11 dígitos
	return re.MatchString(phone)
}

// Função para criar hash da senha
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Criar Usuário com validação
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Dados inválidos"})
	}

	// Remover espaços extras
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Role = strings.TrimSpace(user.Role)
	user.Phone = strings.TrimSpace(user.Phone)
	user.Cpf = strings.TrimSpace(user.Cpf)

	// Validações
	if user.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Nome é obrigatório"})
	}
	if !isValidEmail(user.Email) {
		return c.Status(400).JSON(fiber.Map{"error": "E-mail inválido"})
	}
	if !isValidPhone(user.Phone) {
		return c.Status(400).JSON(fiber.Map{"error": "Telefone inválido"})
	}
	if !isValidCPF(user.Cpf) {
		return c.Status(400).JSON(fiber.Map{"error": "CPF inválido"})
	}
	if user.PasswordHash == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Senha é obrigatória"})
	}
	if user.Role != "cliente" && user.Role != "diarista" {
		return c.Status(400).JSON(fiber.Map{"error": "Papel deve ser 'cliente' ou 'diarista'"})
	}

	// Verificar se já existe um usuário com o mesmo email, telefone ou CPF
	existingUser := models.User{}
	if err := config.DB.Where("email = ? OR phone = ? OR cpf = ?", user.Email, user.Phone, user.Cpf).First(&existingUser).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{"error": "E-mail, telefone ou CPF já cadastrado"})
	}

	// Criar hash da senha
	hashedPassword, err := hashPassword(user.PasswordHash)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao criar senha"})
	}
	user.PasswordHash = hashedPassword

	// Criar usuário no banco
	config.DB.Create(&user)

	// Retornar usuário sem expor a senha
	user.PasswordHash = ""
	return c.JSON(fiber.Map{"id": user.ID})
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
