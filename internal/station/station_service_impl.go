package station

type StationServiceImpl struct{}

func (s *StationServiceImpl) GetAllStations() ([]Station, error) {

	stationList := []Station{
		{
			Latitude:      55.7558,
			Longitude:     37.6176,
			Address:       "Red Square, Moscow",
			StationID:     "station001",
			StationName:   "Moscow Charge Hub",
			StationType:   "Fast",
			AvailableAt:   nil,
			Connectors:    []Connector{},
			PowerKW:       50,
			Price:         500.0,
			PriceUnit:     "kWh",
			PriceCurrency: "RUB",
		},
		{
			Latitude:      59.9343,
			Longitude:     30.3351,
			Address:       "Nevsky Prospect, SPB",
			StationID:     "station002",
			StationName:   "SPB EV Station",
			StationType:   "Standard",
			AvailableAt:   nil, // не доступна в данный момент
			Connectors:    []Connector{},
			PowerKW:       7,
			Price:         250.0,
			PriceUnit:     "hour",
			PriceCurrency: "RUB",
		},
	}

	return stationList, nil
}

func (s *StationServiceImpl) GetStationsByLocation(latitude, longitude float64) ([]Station, error) {
	return nil, nil
}

func (s *StationServiceImpl) GetStationByID(stationID string) (Station, error) {
	return Station{}, nil
}

func (s *StationServiceImpl) GetStationsByType(stationType string) ([]Station, error) {
	return nil, nil
}

func (s *StationServiceImpl) GetStationsByConnector(connectorType Connector) ([]Station, error) {
	return nil, nil
}

func (s *StationServiceImpl) GetStationsByPriceRange(minPrice, maxPrice float64) ([]Station, error) {
	return nil, nil
}

func (s *StationServiceImpl) GetStationsByAvailability(availableAt string) ([]Station, error) {
	return nil, nil
}

func (s *StationServiceImpl) GetStationsByPowerRange(minPower, maxPower float64) ([]Station, error) {
	return nil, nil
}

func (s *StationServiceImpl) GetStationsByPriceUnit(priceUnit string) ([]Station, error) {
	return nil, nil
}

func (s *StationServiceImpl) GetStationsByPriceCurrency(priceCurrency string) ([]Station, error) {
	return nil, nil
}

func (s *StationServiceImpl) GetStationsByAddress(address string) ([]Station, error) {
	return nil, nil
}

func (s *StationServiceImpl) GetStationsByName(stationName string) ([]Station, error) {
	return nil, nil
}
