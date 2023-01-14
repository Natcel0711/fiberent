package routes

import (
	"fiberent/app/controllers"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	//antes del handle los routes empizan asi /api/v1
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/user/:username", controllers.GetUser)
	route.Get("/users", controllers.GetAllUsers)
}
