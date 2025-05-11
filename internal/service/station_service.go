package service

import (
	"main/internal/domain/model"
)

type StationService interface {
	GetAllStations() ([]model.Station, error)
	GetStationsByLocation(latitude, longitude float64) ([]model.Station, error)
	GetStationByID(stationID string) (model.Station, error)
	GetStationsByType(stationType string) ([]model.Station, error)
	GetStationsByConnector(connectorType model.Connector) ([]model.Station, error)
	GetStationsByPriceRange(minPrice, maxPrice float64) ([]model.Station, error)
	GetStationsByAvailability(availableAt string) ([]model.Station, error)
	GetStationsByPowerRange(minPower, maxPower float64) ([]model.Station, error)
	GetStationsByPriceUnit(priceUnit string) ([]model.Station, error)
	GetStationsByPriceCurrency(priceCurrency string) ([]model.Station, error)
	GetStationsByAddress(address string) ([]model.Station, error)
	GetStationsByName(stationName string) ([]model.Station, error)
	CreateStation(station model.Station) (string, error)
	UpdateStation(station model.Station) error
	DeleteStation(stationID string) (bool, error)
	FindNearestStations(lat, lon float64, limit int) ([]model.StationWithDistance, error)
	GetStationsByModerationStatus(status string) ([]model.Station, error)
}
