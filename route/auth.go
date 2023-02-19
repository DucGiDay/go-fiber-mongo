package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hiiamtrong/go-fiber-restapi/controller"
)

func UserRouteAuth(route fiber.Router) {
	route.Post("/login", controller.Login)
	route.Post("/register", controller.Register)
}
