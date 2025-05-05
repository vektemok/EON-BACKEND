package station

import (
	"main/internal/storage/sql/gen"
)

type StationService interface {
	GetAllStations() ([]gen.Station, error)
	GetStationsByLocation(latitude, longitude float64) ([]gen.Station, error)
	GetStationByID(stationID string) (gen.Station, error)
	GetStationsByType(stationType string) ([]gen.Station, error)
	GetStationsByConnector(connectorType Connector) ([]gen.Station, error)
	GetStationsByPriceRange(minPrice, maxPrice float64) ([]gen.Station, error)
	GetStationsByAvailability(availableAt string) ([]gen.Station, error)
	GetStationsByPowerRange(minPower, maxPower float64) ([]gen.Station, error)
	GetStationsByPriceUnit(priceUnit string) ([]gen.Station, error)
	GetStationsByPriceCurrency(priceCurrency string) ([]gen.Station, error)
	GetStationsByAddress(address string) ([]gen.Station, error)
	GetStationsByName(stationName string) ([]gen.Station, error)
	CreateStation(station gen.Station) (string, error)
	UpdateStation(stationID string, station gen.Station) error
	DeleteStation(stationID string) (bool, error)
	FindNearestStations(lat, lon float64, limit int) ([]StationWithDistance, error)
}
