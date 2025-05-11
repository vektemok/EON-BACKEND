package model

import "time"

type Station struct {
	StationID        string    `json:"station_id"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
	Address          string    `json:"address"`
	StationName      string    `json:"station_name"`
	StationType      string    `json:"station_type"`
	AvailableAt      time.Time `json:"available_at"`
	Connectors       []string  `json:"connectors"`
	PowerKw          float64   `json:"power_kw"`
	Price            float64   `json:"price"`
	PriceUnit        string    `json:"price_unit"`
	PriceCurrency    string    `json:"price_currency"`
	ModerationStatus string    `json:"moderation_status"`
}
