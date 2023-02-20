package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(c *fiber.Ctx) error {
    // Lấy JWT từ header "Authorization"
    authHeader := c.Get("Authorization")
    if authHeader == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Authorization token missing",
        })
    }

    // Tách JWT từ chuỗi header "Bearer <token>"
    token := strings.Split(authHeader, " ")[1]

    // Giải mã JWT để lấy thông tin người dùng
    claims := jwt.MapClaims{}
    _, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte("secret"), nil // Giải mã bằng key "secret"
    })
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Invalid authorization token",
        })
    }

    // Lưu thông tin người dùng vào context
    c.Locals("user", claims)

    return c.Next()
}

