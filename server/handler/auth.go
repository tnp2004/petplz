package handler

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tnp2004/petplz/service"
)

var SecretKey = os.Getenv("SECRET_KEY")

func (h accountHandler) Login(c *fiber.Ctx) error {

	var loginData service.LoginRequired
	if err := c.BodyParser(&loginData); err != nil {
		return err
	}

	account, err := h.accountService.Login(loginData.Email, loginData.Password)
	if err != nil {
		return err
	}

	// JWT
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    account.AccountID.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return err
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(account)
}
