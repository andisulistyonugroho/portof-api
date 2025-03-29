package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"portof-api/model"

	"github.com/gofiber/fiber/v3"
)


type LoginRequest struct {
	Username string `json:"username"`
	Passw string `json:"passw"`
}

func (cdh *DuckHandler) Login(c fiber.Ctx) error {

	req := new(LoginRequest)

	if err := c.Bind().JSON(req); err != nil {
		return fiber.ErrBadRequest
	}

	// password hashing
	s := req.Passw
	hasher := sha256.New()
	hasher.Write([]byte(s))
	hashersum := hasher.Sum(nil)
	hashedpwd := hex.EncodeToString(hashersum)

	// model.User
	var u model.User
	row := cdh.DB.QueryRow(`SELECT id, username FROM users WHERE is_verified = true AND username = ? AND passw = ?`, req.Username, hashedpwd)
	
	err := row.Scan(&u.ID, &u.Username)
	if err != nil {
		fmt.Println("scan error:", err)
		return err
	}


	return c.JSON(fiber.Map{
		"id": u.ID,
		"username": u.Username,
	})
}