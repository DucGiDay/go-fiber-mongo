package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hiiamtrong/go-fiber-restapi/controller"
)

func UserRoute(route fiber.Router) {
	route.Get("/", controller.GetAllUsers)
	route.Post("/", controller.CreateUser)
	route.Get("/:userId", controller.GetUser)
	route.Put("/:userId", controller.UpdateUser)
}
