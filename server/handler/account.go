package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tnp2004/petplz/service"
)

type accountHandler struct {
	accountService service.AccountService
}

func NewAccoutHandler(accountService service.AccountService) accountHandler {
	return accountHandler{accountService}
}

func (h accountHandler) Greeting(c *fiber.Ctx) error {
	return c.SendString("hi there!")
}

func (h accountHandler) GetAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	account, err := h.accountService.GetUserAccount(id)
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(account)
}

func (h accountHandler) CreateAccount(c *fiber.Ctx) error {
	var account service.NewAccount

	if err := c.BodyParser(&account); err != nil {
		return c.JSON(err)
	}

	err := h.accountService.NewAccount(account)
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(fiber.Map{
		"message": "your account has been created",
	})
}
