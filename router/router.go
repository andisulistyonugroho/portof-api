package router

import (
	"database/sql"
	"portof-api/controller"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(r *fiber.App, db *sql.DB) {

	cdh := &controller.DuckHandler{DB: db}

	// auth
	auth := r.Group("/user")
	auth.Post("/login", cdh.Login)

	// public page
	r.Get("/pages", cdh.ListPage)

	// admin
	// mabes := r.Group("/mabes")
	// mabesPage := mabes.Group("/pages")
	// mabesPage.Post("/", cdh.ListPage)
}
