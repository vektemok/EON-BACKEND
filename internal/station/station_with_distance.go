package station

import "main/internal/storage/sql/gen"

type StationWithDistance struct {
	Station  *gen.Station
	Distance float64
}
