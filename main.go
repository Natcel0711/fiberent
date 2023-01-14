package main

import (
	"fiberent/pkg/configs"
	"fiberent/pkg/routes"
	"fiberent/pkg/utils"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	config := configs.FiberConfig()
	app := fiber.New(config)

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)

	utils.StartServer(app)
}
