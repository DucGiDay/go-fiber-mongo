package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hiiamtrong/go-fiber-restapi/config"
	"github.com/hiiamtrong/go-fiber-restapi/route"
	"github.com/hiiamtrong/go-fiber-restapi/middleware"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	config.ConnectDatabase()

	setupRoutes(app)
	app.Listen(":" + os.Getenv("PORT"))
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")

	route.UserRoute(api.Group("/users", middleware.AuthMiddleware))
	route.UserRouteAuth(api.Group("/auth"))
}
