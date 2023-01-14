package controllers

import (
	"fiberent/app/models"
	"fiberent/ent"
	"fiberent/ent/usuario"
	"fiberent/pkg/repository"
	"fiberent/pkg/utils"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	url, _ := utils.ConnectionURLBuilder("postgres")
	client, err := ent.Open("postgres", url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Something happened in the server",
			"user":    nil,
		})
	}
	users, err := client.Usuario.Query().All(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Something happened retrieving users",
			"user":    nil,
		})
	}
	return c.JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"users":   users,
	})
}

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

func CreateUser(c *fiber.Ctx) error {
	now := time.Now().Unix()
	log.Println(now)
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	expires := claims.Expires

	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	credential := claims.Credentials[repository.UserCreateCredential]

	if !credential {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "permission denied, check credential of your token",
		})
	}

	user := &models.User{}

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	url, _ := utils.ConnectionURLBuilder("postgres")
	client, err := ent.Open("postgres", url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Error connecting to db",
		})
	}

	UserCreated, err := client.Usuario.Create().
		SetUsername(user.Username).
		SetPassword(user.PasswordHash).
		SetEmail(user.Email).Save(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "User Created",
		"User":    UserCreated,
	})
}
