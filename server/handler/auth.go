package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tnp2004/petplz/service"
)

func (h accountHandler) Login(c *fiber.Ctx) error {

	var loginData service.LoginRequired
	if err := c.BodyParser(&loginData); err != nil {
		return err
	}

	err := h.accountService.Login(loginData.Email, loginData.Password)
	if err != nil {
		return err
	}

	return c.SendString("you are login successfully")
}
