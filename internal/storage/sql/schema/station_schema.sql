CREATE TABLE stations (
    station_id      TEXT PRIMARY KEY,
    latitude        DOUBLE PRECISION NOT NULL,
    longitude       DOUBLE PRECISION NOT NULL,
    address         TEXT,
    station_name    TEXT,
    station_type    TEXT,
    available_at    TIMESTAMP,
    connectors      TEXT[],  
    power_kw        DOUBLE PRECISION NOT NULL,
    price           DOUBLE PRECISION NOT NULL,
    price_unit      TEXT NOT NULL,
    price_currency  TEXT NOT NULL
);
