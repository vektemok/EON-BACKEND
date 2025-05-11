package model

import (
	"main/internal/storage/sql/gen"

	"github.com/jackc/pgx/v5/pgtype"
)

func MapStationToGenStation(station *Station) *gen.Station {
	return &gen.Station{
		StationID:        station.StationID,
		Latitude:         station.Latitude,
		Longitude:        station.Longitude,
		Address:          pgtype.Text{String: station.Address, Valid: true},
		StationName:      pgtype.Text{String: station.StationName, Valid: true},
		StationType:      pgtype.Text{String: station.StationType, Valid: true},
		AvailableAt:      pgtype.Timestamp{Time: station.AvailableAt, Valid: true},
		Connectors:       station.Connectors,
		PowerKw:          station.PowerKw,
		Price:            station.Price,
		PriceUnit:        station.PriceUnit,
		PriceCurrency:    station.PriceCurrency,
		ModerationStatus: station.ModerationStatus,
	}

}

func MapGenStationToStation(genStation *gen.Station) *Station {
	return &Station{
		StationID:        genStation.StationID,
		Latitude:         genStation.Latitude,
		Longitude:        genStation.Longitude,
		Address:          genStation.Address.String,
		StationName:      genStation.StationName.String,
		StationType:      genStation.StationType.String,
		AvailableAt:      genStation.AvailableAt.Time,
		Connectors:       genStation.Connectors,
		PowerKw:          genStation.PowerKw,
		Price:            genStation.Price,
		PriceUnit:        genStation.PriceUnit,
		PriceCurrency:    genStation.PriceCurrency,
		ModerationStatus: genStation.ModerationStatus,
	}
}

func MapStationToCreateStationParams(station *Station) *gen.CreateStationParams {
	return &gen.CreateStationParams{
		StationID:        station.StationID,
		Latitude:         station.Latitude,
		Longitude:        station.Longitude,
		Address:          pgtype.Text{String: station.Address, Valid: true},
		StationName:      pgtype.Text{String: station.StationName, Valid: true},
		StationType:      pgtype.Text{String: station.StationType, Valid: true},
		AvailableAt:      pgtype.Timestamp{Time: station.AvailableAt, Valid: true},
		Connectors:       station.Connectors,
		PowerKw:          station.PowerKw,
		Price:            station.Price,
		PriceUnit:        station.PriceUnit,
		PriceCurrency:    station.PriceCurrency,
		ModerationStatus: station.ModerationStatus,
	}
}

func MapStationToUpdateStationParams(station *Station) *gen.UpdateStationParams {
	return &gen.UpdateStationParams{
		StationID:     station.StationID,
		Latitude:      station.Latitude,
		Longitude:     station.Longitude,
		Address:       pgtype.Text{String: station.Address, Valid: true},
		StationName:   pgtype.Text{String: station.StationName, Valid: true},
		StationType:   pgtype.Text{String: station.StationType, Valid: true},
		AvailableAt:   pgtype.Timestamp{Time: station.AvailableAt, Valid: true},
		Connectors:    station.Connectors,
		PowerKw:       station.PowerKw,
		Price:         station.Price,
		PriceUnit:     station.PriceUnit,
		PriceCurrency: station.PriceCurrency,
	}
}
