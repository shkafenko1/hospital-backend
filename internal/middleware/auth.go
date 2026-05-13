package middleware

import (
	"hospital-backend/internal/service"
	"strings"

	"github.com/gofiber/fiber/v3"
)

type AuthMiddleware struct {
	service *service.UserService
}

func NewAuthMiddleware(service *service.UserService) *AuthMiddleware {
	return &AuthMiddleware{service: service}
}

func (m *AuthMiddleware) Authenticate(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing authorization header",
		})
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid authorization header format",
		})
	}

	claims, err := m.service.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	c.Locals("user_id", claims["user_id"])
	c.Locals("username", claims["username"])
	c.Locals("role", claims["role"])

	return c.Next()
}

func (m *AuthMiddleware) RequireRole(role string) fiber.Handler {
	return func(c fiber.Ctx) error {
		userRole := c.Locals("role")
		if userRole != role {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Insufficient permissions",
			})
		}
		return c.Next()
	}
}

func (m *AuthMiddleware) RequireAdmin() fiber.Handler {
	return m.RequireRole("admin")
}
