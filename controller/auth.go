package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)


type loginRequest struct {
	Username string `json:"username"`
	Passw string `json:"passw"`
}

func (cdh *DuckHandler) Login(c fiber.Ctx) error {

	req := new(loginRequest)

	if err := c.Bind().Body(req); err != nil {
		return fiber.ErrBadRequest
	}

	// model.User
	// u := new(model.User)
	// var u model.User
	// var (
	// 	ID int
	// 	Username string
	// )

	// row := datasources.DB.QueryRow(`SELECT id, username FROM users WHERE username = ? and passw = ?`, req.Username, req.Passw)
	// row := datasources.DB.QueryRow(`SELECT id, username FROM users`)
	// err := row.Scan(&ID, &Username)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return err
	// }

	// datasources.DB.QueryRow("SELECT * FROM users")
	// datasources.Duckconnect()
	fmt.Println("DATASOURCE DB:",cdh.DB)
	defer cdh.DB.Close()


	return c.JSON(fiber.Map{
		"OK": "ID",
		"Username": "Username",
	})
}