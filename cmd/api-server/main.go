package main

import (
	"context"
	"fmt"
	// "main/internal/auth"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func main() {
	urlExample := "postgres://postgres:1234@localhost:4321/postgres"

	conn, err := pgx.Connect(context.Background(), urlExample)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())
	app := fiber.New()

	var dbVersion string
	err = conn.QueryRow(context.Background(), "SELECT version()").Scan(&dbVersion)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute query: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("PostgreSQL version:", dbVersion)

	// app.Get("/auth:name?", auth.SignInByPhoneNumber)

	app.Listen(":3000")

}
