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
		Issuer:    account.AccountID.Hex(),
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

func (h accountHandler) Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "logout successfully",
	})
}

func (h accountHandler) LoginAuth(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		// not login
		return err
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	_, err = h.accountService.GetUserAccount(claims.Issuer)
	if err != nil {
		return err
	}

	return c.Next()
}
