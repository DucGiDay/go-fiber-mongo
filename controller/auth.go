package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hiiamtrong/go-fiber-restapi/model"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Tạo một secret key để ký JWT
var secretKey = []byte("secret")

//Login Controller
func Login(c *fiber.Ctx) error {
	// Lấy thông tin người dùng từ request body
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	// Tìm kiếm người dùng trong db
	user, err := model.CheckInvalidCredentials(user.Username, user.Password)

	// Nếu không tìm thấy người dùng
	if err != nil {
		return c.Status(401).SendString("Invalid credentials")
	}

	// Nếu tìm thấy người dùng, tạo JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return c.Status(500).SendString("Failed to create token")
	}

	// Trả về JWT token
	return c.JSON(tokenString)

}

//Register Controller
func Register(c *fiber.Ctx) error {
	// Lấy thông tin người dùng từ request body
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	// Kiểm tra xem người dùng đã tồn tại trong danh sách hay chưa
	_, err := model.CheckUserAlreadyExists(user.Username)

	// Nếu không tìm thấy người dùng 
	if err != nil  {
		user.ID = primitive.NewObjectID()
		user, err := model.CreateUser(user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": fiber.ErrInternalServerError.Message,
				"error":   err,
			})
		}
		return c.JSON(user)
	}
	// Nếu tìm thấy người dùng báo lỗi
	return c.Status(400).SendString("Username already exists")
}