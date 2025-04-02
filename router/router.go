package router

import (
	"database/sql"
	"portof-api/controller"
	"portof-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(r *fiber.App, db *sql.DB) {

	cdh := &controller.DuckHandler{DB: db}

	// auth
	auth := r.Group("/user")
	auth.Post("/login", cdh.Login)

	// public page
	r.Get("/pages", cdh.ListPage)

	// admin
	mabes := r.Group("/mabes")
	mabes.Use(middleware.JWTProtected)
	mabesPage := mabes.Group("/pages")
	mabesPage.Post("/", cdh.CreatePage)
}
