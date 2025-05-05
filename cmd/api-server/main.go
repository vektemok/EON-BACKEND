package main

import (
	"main/internal/config"
	"main/internal/lib/logger"
	"main/internal/storage/postgres"
	"strconv"

	// "main/internal/storage/sql/gen"
	db "main/internal/storage/sql/gen"
	// "strconv"
	"time"

	"main/internal/station"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
)

type StationDTO struct {
	StationID   string    `json:"station_id"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Address     string    `json:"address"`
	StationType string    `json:"station_type"`
	StationName string    `json:"station_name"`
	AvailableAt time.Time `json:"available_at"`
	PowerKw     float64   `json:"power_kw"`

	Price         float64  `json:"price"`
	PriceUnit     string   `json:"price_unit"`
	PriceCurrency string   `json:"price_currency"`
	Connectors    []string `json:"connectors"`
}

func main() {

	cfg := config.LoadConfig()

	log := logger.SetupLogger(cfg.Env)

	dbPool := postgres.ConnectDB(cfg, log)

	queries := db.New(dbPool)

	stationService := station.NewStationService(queries)

	app := fiber.New()

	app.Get("/nearest", func(c *fiber.Ctx) error {
		lat, err1 := strconv.ParseFloat(c.Query("lat"), 64)
		lon, err2 := strconv.ParseFloat(c.Query("lon"), 64)
		limit, err3 := strconv.Atoi(c.Query("limit", "3"))

		if err1 != nil || err2 != nil || err3 != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid query parameters",
			})
		}

		stationsWithDistance, err := stationService.FindNearestStations(lat, lon, limit)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(stationsWithDistance)

	})
	app.Get("/stations", func(c *fiber.Ctx) error {
		stations, err := stationService.GetAllStations()
		log.Sugar().Errorf("Failed to get stations: %v", err)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err,
			})
		}

		return c.JSON(stations)
	})

	app.Post("/stations", func(c *fiber.Ctx) error {
		var dto StationDTO

		if err := c.BodyParser(&dto); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		station := db.Station{
			StationID:     dto.StationID,
			Latitude:      dto.Latitude,
			Longitude:     dto.Longitude,
			Address:       pgtype.Text{String: dto.Address, Valid: true},
			StationName:   pgtype.Text{String: dto.StationName, Valid: true},
			StationType:   pgtype.Text{String: dto.StationType, Valid: true},
			AvailableAt:   pgtype.Timestamp{Time: dto.AvailableAt, Valid: true},
			Connectors:    dto.Connectors,
			PowerKw:       dto.PowerKw,
			Price:         dto.Price,
			PriceUnit:     dto.PriceUnit,
			PriceCurrency: dto.PriceCurrency,
		}

		createdStation, err := stationService.CreateStation(station)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(createdStation)
	})

	err := app.Listen(cfg.Address)
	if err != nil {
		log.Sugar().Fatalf("Failed to start server: %v", err)
	}

	dbPool.Close()
}
