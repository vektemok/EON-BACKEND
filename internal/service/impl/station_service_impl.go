package service

import (
	"context"
	"sort"

	"main/internal/lib/haversine"
	"main/internal/storage/sql/gen"

	"go.uber.org/zap"

	"main/internal/domain/model"
)

func NewStationService(queries *gen.Queries) *StationServiceImpl {
	return &StationServiceImpl{queries: queries}
}

type StationServiceImpl struct {
	queries *gen.Queries
}

func (s *StationServiceImpl) GetAllStations() ([]model.Station, error) {
	ctx := context.Background()
	genStations, err := s.queries.ListStations(ctx) // returned []gen.Station

	if err != nil {
		return nil, err
	}

	if genStations == nil {
		return []model.Station{}, nil
	}

	stations := make([]model.Station, len(genStations))
	for i, genStation := range genStations {
		stations[i] = *model.MapGenStationToStation(&genStation)
	}

	return stations, nil
}

func (s *StationServiceImpl) FindNearestStations(lat, lon float64, limit int) ([]model.StationWithDistance, error) {
	ctx := context.Background()
	genStations, err := s.queries.ListStations(ctx)
	if err != nil {
		return nil, err
	}

	stations := make([]model.Station, len(genStations))
	for i, genStation := range genStations {
		stations[i] = *model.MapGenStationToStation(&genStation)
	}

	distances := make([]model.StationWithDistance, 0, len(stations))

	for i := range stations {
		dist := haversine.Haversine(lat, lon, stations[i].Latitude, stations[i].Longitude)
		distances = append(distances, model.StationWithDistance{
			Station:  model.MapStationToGenStation(&stations[i]),
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

func (s *StationServiceImpl) GetStationsByModerationStatus(status string) ([]model.Station, error) {
	ctx := context.Background()
	genStations, err := s.queries.GetStationsByModerationStatus(ctx, status) // returned []gen.Station

	if err != nil {
		return nil, err
	}

	if genStations == nil {
		return []model.Station{}, nil
	}

	stations := make([]model.Station, len(genStations))
	for i, genStation := range genStations {
		stations[i] = *model.MapGenStationToStation(&genStation)
	}

	zap.L().Info("GetStationsByModerationStatus", zap.Reflect("stations", stations))

	return stations, nil
}

func (s *StationServiceImpl) CreateStation(station model.Station) (model.Station, error) {
	ctx := context.Background()

	params := model.MapStationToCreateStationParams(&station) // Map Station to gen.Station
	createdGenStation, err := s.queries.CreateStation(ctx, *params)

	if err != nil {
		return model.Station{}, err
	}

	createdStation := model.MapGenStationToStation(&createdGenStation)
	return *createdStation, nil
}

func (s *StationServiceImpl) UpdateStation(stationID string, station model.Station) error {
	ctx := context.Background()

	params := model.MapStationToUpdateStationParams(&station)

	_, err := s.queries.UpdateStation(ctx, *params)

	return err
}

func (s *StationServiceImpl) DeleteStation(stationID string) (bool, error) {
	ctx := context.Background()
	err := s.queries.DeleteStation(ctx, stationID)

	if err != nil {
		return false, err
	}

	return true, nil
}
