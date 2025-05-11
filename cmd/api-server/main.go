package main

import (
	// "fmt"
	// "encoding/json"
	// "fmt"
	"main/internal/config"
	"main/internal/domain/model"
	"main/internal/service/impl"

	// "main/internal/lib/api/response"
	l "log"
	"main/internal/lib/api/response"
	"main/internal/lib/logger"
	"main/internal/storage/postgres"
	"strconv"

	db "main/internal/storage/sql/gen"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	cfg := config.LoadConfig()

	log := logger.SetupLogger(cfg.Env)

	dbPool := postgres.ConnectDB(cfg, log)

	queries := db.New(dbPool)

	app := fiber.New()

	app.Use(cors.New())

	stationService := service.NewStationService(queries)

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

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
		var station model.Station

		if err := c.BodyParser(&station); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		createdStation, err := stationService.CreateStation(station)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(createdStation)
	})

	app.Put("/stations/:station_id", func(c *fiber.Ctx) error {

		stationID := c.Params("station_id")

		if stationID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(response.Error("station_id is required"))
		}

		var station model.Station

		err := stationService.UpdateStation(stationID, station)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
		}

		return c.Status(fiber.StatusOK).JSON(response.OK())

	})

	app.Get("/stations/:station_id", func(c *fiber.Ctx) error {
		stationID := c.Params("station_id")

		if stationID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(response.Error("station_id is required"))
		}

		stations, err := stationService.GetStationsByModerationStatus(stationID)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.Error("error fetch data"))
		}

		return c.Status(fiber.StatusOK).JSON(stations)
	})

	l.Fatal(app.Listen(cfg.Address))

	dbPool.Close()
}
