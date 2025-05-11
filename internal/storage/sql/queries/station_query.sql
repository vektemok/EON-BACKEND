-- name: CreateStation :one
INSERT INTO stations (
    station_id, latitude, longitude, address, station_name,
    station_type, available_at, connectors, power_kw,
    price, price_unit, price_currency, moderation_status
) VALUES (
    $1, $2, $3, $4, $5,
    $6, $7, $8, $9,
    $10, $11, $12, $13
)
RETURNING *;

-- name: GetStation :one
SELECT * FROM stations
WHERE station_id = $1;

-- name: ListStations :many
SELECT * FROM stations
ORDER BY station_name;

-- name: UpdateStation :one
UPDATE stations
SET
    latitude = $2,
    longitude = $3,
    address = $4,
    station_name = $5,
    station_type = $6,
    available_at = $7,
    connectors = $8,
    power_kw = $9,
    price = $10,
    price_unit = $11,
    price_currency = $12
WHERE station_id = $1
RETURNING *;

-- name: DeleteStation :exec
DELETE FROM stations
WHERE station_id = $1;

-- name: GetStationsByModerationStatus :many
SELECT * FROM stations
WHERE moderation_status = $1;