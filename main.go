package main

import (
	"context"
	"fiberent/ent"
	"fiberent/ent/usuario"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("ay")
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres password=Tripleh1 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer client.Close()
	user, _ := QueryUser(context.Background(), client)
	fmt.Println(user)
}

func CreateUsuario(ctx context.Context, client *ent.Client) (*ent.Usuario, error) {
	u, err := client.Usuario.Create().
		SetEmail("natcelnieves@gmail.com").SetUsername("nievesarce").SetPassword("Tripleh1").Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.Usuario, error) {
	u, err := client.Usuario.Query().Where(usuario.Username("nievesarce")).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %v ", err)
	}
	log.Println("user found: ", u)
	return u, nil
}
