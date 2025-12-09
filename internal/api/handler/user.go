package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/AliasgharHeidari/test-gorm/internal/model"
	"github.com/AliasgharHeidari/test-gorm/internal/security"
	"github.com/AliasgharHeidari/test-gorm/internal/service"
)

func CreateUser(c *fiber.Ctx) error {
	var inputUser model.User

	err := c.BodyParser(&inputUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := service.CreateUser(inputUser); errors.Is(err, security.ErrEmptyPassword) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "password can not be empty",
		})
	} else if errors.Is(err, security.ErrLengthPassword) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "password must be at least 8 characters",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal error : failed to add user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user has been added successfuly",
		"user":    inputUser.UserName,
	})

}
