package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hiiamtrong/go-fiber-restapi/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllUsers(c *fiber.Ctx) error {
	users, err := model.GetAllUsers()
	if err != nil {
		return err
	}

	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	userId, err := primitive.ObjectIDFromHex(c.Params("userId"))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrBadRequest.Message,
			"error":   err,
		})
	}
	user, err := model.GetUser(userId)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrBadRequest.Message,
			"error":   err,
		})
	}
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": fiber.ErrBadRequest.Message,
				"error":   err,
			})
		}
	}

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

func UpdateUser(c *fiber.Ctx) error {
	var user model.User
	userId, err := primitive.ObjectIDFromHex(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrNotFound.Message,
			"error":   err,
		})
	}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user, err = model.UpdateUser(userId, user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrInternalServerError.Message,
			"error":   err,
		})
	}

	return c.JSON(user)
}
