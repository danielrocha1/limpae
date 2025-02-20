package routes

import (
	"limpae/go/src/config"
	"limpae/go/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"os"
)

type CheckoutRequest struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	Street    string  `json:"street"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	Zipcode   string  `json:"zip"`
	Plan      string  `json:"plan"`
	StripeToken string `json:"stripeToken"`
}

func CheckoutHandler(c *fiber.Ctx) error {
	var request CheckoutRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Processar pagamento na Stripe
	stripe.Key = os.Getenv("STRIPE_SECRET")
	chargeParams := &stripe.ChargeParams{
		Amount:   stripe.Int64(getPlanPrice(request.Plan)),
		Currency: stripe.String("brl"),
		Source:   &stripe.SourceParams{Token: stripe.String(request.StripeToken)},
	}
	_, err := charge.New(chargeParams)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Payment failed"})
	}

	// Criar usuário no banco
	user := models.User{
		Name:  request.Name,
		Email: request.Email,
		Phone: request.Phone,
		Role:  "cliente",
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	// Criar endereço
	address := models.Address{
		UserID:  user.ID,
		Street:  request.Street,
		City:    request.City,
		State:   request.State,
		Zipcode: request.Zipcode,
	}
	if err := config.DB.Create(&address).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save address"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Payment successful, user created!"})
}

func getPlanPrice(plan string) int64 {
	switch plan {
	case "basic":
		return 1990
	case "premium":
		return 4990
	default:
		return 0
	}
}
