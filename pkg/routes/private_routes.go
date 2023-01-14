package routes

import (
	"fiberent/app/controllers"
	"fiberent/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(app *fiber.App) {
	routes := app.Group("/api/v1")

	routes.Post("/user", middleware.JWTProtected(), controllers.CreateUser)

	//client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres password=Tripleh1 dbname=postgres sslmode=disable")
	/*if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer client.Close()
	*/
}

/*
func CreateUser(ctx context.Context, client *ent.Client) (*ent.Usuario, error) {
	u, err := client.Usuario.Create().
		SetEmail("natcelnieves@gmail.com").SetUsername("nievesarce").SetPassword("Tripleh1").Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client, username string) (*ent.Usuario, error) {
	user, err := client.Usuario.Query().Where(usuario.Username(username)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %v", err)
	}
	return user, nil
}
*/
