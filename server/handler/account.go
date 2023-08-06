package handler

import (
	"net/http"

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
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(account)
}

func (h accountHandler) CreateAccount(c *fiber.Ctx) error {
	var account service.NewAccount

	if err := c.BodyParser(&account); err != nil {
		return c.JSON(err)
	}

	err := h.accountService.Register(account)
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(fiber.Map{
		"message": "your account has been created",
	})
}

func (h accountHandler) ValidateRegisterForm(c *fiber.Ctx) error {
	value := c.Queries()

	if value["email"] != "" {
		if err := h.accountService.ValidateEmail(value["email"]); err != nil {
			c.Status(http.StatusConflict)
			return c.JSON(fiber.Map{
				"message": "email is used",
			})
		}

		return c.JSON(fiber.Map{
			"message": "email is unused",
		})

	} else if value["username"] != "" {
		if err := h.accountService.ValidateUsername(value["username"]); err != nil {
			c.Status(http.StatusConflict)
			return c.JSON(fiber.Map{
				"message": "username is used",
			})
		}

		return c.JSON(fiber.Map{
			"message": "username is unused",
		})
	}

	c.Status(http.StatusBadRequest)
	return c.JSON(fiber.Map{
		"message": "invalid data",
	})
}
