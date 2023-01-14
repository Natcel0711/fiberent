package controllers

import (
	"fiberent/ent"
	"fiberent/ent/usuario"
	"fiberent/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func GetUser(x *fiber.Ctx) error {
	username := x.Params("username")
	url, _ := utils.ConnectionURLBuilder("postgres")
	client, err := ent.Open("postgres", url)
	if err != nil {
		return x.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "something bad happened in the server",
			"user":    nil,
		})
	}
	user, err := client.Usuario.Query().Where(usuario.Username(username)).Only(x.Context())
	if err != nil {
		return x.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "user with given username not found",
			"user":    nil,
		})
	}

	return x.JSON(fiber.Map{
		"error":    false,
		"message":  nil,
		"username": user.Username,
		"id":       user.ID,
		"email":    user.Email,
	})
}
