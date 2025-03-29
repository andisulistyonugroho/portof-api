package main

import (
	"portof-api/datasources"
	"portof-api/router"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	duck := datasources.Duckconnect()
	
	fib := fiber.New()
	fib.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods: []string{"POST", "OPTIONS"},
	}))

	fib.Get("/", func(c fiber.Ctx) error {
		return c.SendString("hello world")
	})

	router.SetupRoutes(fib, duck)

	fib.Listen(":3000")
}

