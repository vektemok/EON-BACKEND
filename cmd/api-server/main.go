package main

import (
	"context"
	"fmt"
	"main/internal/station"

	// "main/internal/auth"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func main() {

	stationService := &station.StationServiceImpl{}

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



	app.Get("/stations", func(c *fiber.Ctx) error {
		stations, err := stationService.GetAllStations()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to retrieve stations",
			})
		}

		return c.JSON(stations)
	})

	app.Listen(":3000")

}
