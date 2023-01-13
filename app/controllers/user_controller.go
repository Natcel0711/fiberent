package controllers

import (
	"fiberent/ent"
	"fiberent/ent/usuario"

	"github.com/gofiber/fiber/v2"
)

func GetUser(x *fiber.Ctx) error {
	username := x.Params("username")
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres password=Tripleh1 dbname=postgres sslmode=disable")
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
