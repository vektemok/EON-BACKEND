package station

import (
	"context"

	"sort"

	"main/internal/lib/haversine"
	"main/internal/storage/sql/gen"
)

//	stations := []gen.Station{
//		{
//			StationID:     "station001",
//			Latitude:      55.7558,
//			Longitude:     37.6176,
//			Address:       sql.NullString{String: "Red Square, Moscow", Valid: true},
//			StationName:   sql.NullString{String: "Moscow Charge Hub", Valid: true},
//			StationType:   sql.NullString{String: "Fast", Valid: true},
//			AvailableAt:   sql.NullTime{Valid: false},
//			Connectors:    []string{"Type2", "CHAdeMO"},
//			PowerKw:       50.0,
//			Price:         500.0,
//			PriceUnit:     "kWh",
//			PriceCurrency: "RUB",
//		},
//		{
//			StationID:     "station002",
//			Latitude:      59.9343,
//			Longitude:     30.3351,
//			Address:       sql.NullString{String: "Nevsky Prospect, Saint Petersburg", Valid: true},
//			StationName:   sql.NullString{String: "SPB EV Station", Valid: true},
//			StationType:   sql.NullString{String: "Standard", Valid: true},
//			AvailableAt:   sql.NullTime{Valid: false},
//			Connectors:    []string{"Type2"},
//			PowerKw:       22.0,
//			Price:         300.0,
//			PriceUnit:     "kWh",
//			PriceCurrency: "RUB",
//		},
//		{
//			StationID:     "station003",
//			Latitude:      56.8389,
//			Longitude:     60.6057,
//			Address:       sql.NullString{String: "Lenina Ave, Yekaterinburg", Valid: true},
//			StationName:   sql.NullString{String: "Ural EV Station", Valid: true},
//			StationType:   sql.NullString{String: "Fast", Valid: true},
//			AvailableAt:   sql.NullTime{Time: time.Now(), Valid: true}, // пример с временем
//			Connectors:    []string{"CCS", "CHAdeMO"},
//			PowerKw:       100.0,
//			Price:         600.0,
//			PriceUnit:     "kWh",
//			PriceCurrency: "RUB",
//		},
//		{
//			StationID:     "station004",
//			Latitude:      43.1056,
//			Longitude:     131.8735,
//			Address:       sql.NullString{String: "Central Square, Vladivostok", Valid: true},
//			StationName:   sql.NullString{String: "Far East EV Stop", Valid: true},
//			StationType:   sql.NullString{String: "Standard", Valid: true},
//			AvailableAt:   sql.NullTime{Valid: false},
//			Connectors:    []string{"Type1"},
//			PowerKw:       11.0,
//			Price:         200.0,
//			PriceUnit:     "hour",
//			PriceCurrency: "RUB",
//		},
//	}
func NewStationService(queries *gen.Queries) *StationServiceImpl {
	return &StationServiceImpl{queries: queries}
}

type StationServiceImpl struct {
	queries *gen.Queries
}

func (s *StationServiceImpl) GetAllStations() ([]gen.Station, error) {
	ctx := context.Background()
	stations, err := s.queries.ListStations(ctx)

	if err != nil {
		return nil, err
	}

	return stations, nil
}

func (s *StationServiceImpl) FindNearestStations(lat, lon float64, limit int) ([]StationWithDistance, error) {
	ctx := context.Background()
	stations, err := s.queries.ListStations(ctx)
	if err != nil {
		return nil, err
	}

	distances := make([]StationWithDistance, 0, len(stations))

	for i := range stations {
		dist := haversine.Haversine(lat, lon, stations[i].Latitude, stations[i].Longitude)
		distances = append(distances, StationWithDistance{
			Station:  &stations[i],
			Distance: dist,
		})
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	if limit > len(distances) {
		limit = len(distances)
	}

	return distances[:limit], nil
}

// Other methods need to be updated similarly. Example:

func (s *StationServiceImpl) GetStationsByLocation(latitude, longitude float64) ([]gen.Station, error) {
	// Example return for location-based stations (you would query from gen in real case)
	return nil, nil
}

func (s *StationServiceImpl) GetStationByID(stationID string) (gen.Station, error) {
	// Example return for station by ID
	return gen.Station{}, nil
}

func (s *StationServiceImpl) GetStationsByType(stationType string) ([]gen.Station, error) {
	return nil, nil
}

func (s *StationServiceImpl) GetStationsByConnector(connectorType string) ([]gen.Station, error) {
	return nil, nil
}

func (s *StationServiceImpl) GetStationsByPriceRange(minPrice, maxPrice float64) ([]gen.Station, error) {
	return nil, nil
}

func (s *StationServiceImpl) GetStationsByAvailability(availableAt string) ([]gen.Station, error) {
	return nil, nil
}

func (s *StationServiceImpl) GetStationsByPowerRange(minPower, maxPower float64) ([]gen.Station, error) {
	return nil, nil
}

func (s *StationServiceImpl) GetStationsByPriceUnit(priceUnit string) ([]gen.Station, error) {
	return nil, nil
}

func (s *StationServiceImpl) GetStationsByPriceCurrency(priceCurrency string) ([]gen.Station, error) {
	return nil, nil
}

func (s *StationServiceImpl) GetStationsByAddress(address string) ([]gen.Station, error) {
	return nil, nil
}

func (s *StationServiceImpl) GetStationsByName(stationName string) ([]gen.Station, error) {
	return nil, nil
}

func (s *StationServiceImpl) CreateStation(station gen.Station) (gen.Station, error) {

	ctx := context.Background()

	params := gen.CreateStationParams(station)

	createdStation, err := s.queries.CreateStation(ctx, params)

	if err != nil {
		return gen.Station{}, err
	}

	return createdStation, nil
}

func (s *StationServiceImpl) UpdateStation(stationID string, station gen.Station) error {
	return nil
}

func (s *StationServiceImpl) DeleteStation(stationID string) (bool, error) {
	return false, nil
}
