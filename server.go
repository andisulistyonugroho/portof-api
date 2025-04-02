package main

import (
	"portof-api/datasources"
	"portof-api/router"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	duck := datasources.Duckconnect()

	fib := fiber.New()
	fib.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: strings.Join([]string{"Origin", "Content-Type", "Accept", "Authorization"}, ","),
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
	}))

	fib.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	router.SetupRoutes(fib, duck)

	fib.Listen(":3000")
}
