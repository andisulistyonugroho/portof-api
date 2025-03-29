package router

import (
	"database/sql"
	"fmt"
	"portof-api/controller"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes (r *fiber.App, db *sql.DB) {

	fmt.Println("SETUP ROUTES:", db)
	cdh := &controller.DuckHandler{DB: db}
	// auth
	auth := r.Group("/user")
	auth.Post("/login", cdh.Login)
}