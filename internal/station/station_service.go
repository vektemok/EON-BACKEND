package station

type StationService interface {
	GetAllStations() ([]Station, error)
	GetStationsByLocation(latitude, longitude float64) ([]Station, error)
	GetStationByID(stationID string) (Station, error)
	GetStationsByType(stationType string) ([]Station, error)
	GetStationsByConnector(connectorType Connector) ([]Station, error)
	GetStationsByPriceRange(minPrice, maxPrice float64) ([]Station, error)
	GetStationsByAvailability(availableAt string) ([]Station, error)
	GetStationsByPowerRange(minPower, maxPower float64) ([]Station, error)
	GetStationsByPriceUnit(priceUnit string) ([]Station, error)
	GetStationsByPriceCurrency(priceCurrency string) ([]Station, error)
	GetStationsByAddress(address string) ([]Station, error)
	GetStationsByName(stationName string) ([]Station, error)
}

