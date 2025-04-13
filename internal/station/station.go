package station

import "time"

type Station struct {
	Latitude       float64     `json:"latitude"`
	Longitude      float64     `json:"longitude"`
	Address        string      `json:"address"`
	StationID      string      `json:"station_id"`
	StationName    string      `json:"station_name"`
	StationType    string      `json:"station_type"`
	AvailableAt    *time.Time  `json:"available_at"`
	Connectors []Connector `json:"connector_types"`
	PowerKW        float64     `json:"power_kw"`
	Price          float64     `json:"price"`
	PriceUnit      string      `json:"price_unit"`
	PriceCurrency  string      `json:"price_currency"`
	
}
