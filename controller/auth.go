package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"portof-api/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Username string `json:"username"`
	Passw    string `json:"passw"`
}

func (cdh *DuckHandler) Login(c *fiber.Ctx) error {

	req := new(LoginRequest)

	if err := c.BodyParser(req); err != nil {
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

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JPORTOFW_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"token":    t,
		"id":       u.ID,
		"username": u.Username,
	})

}
