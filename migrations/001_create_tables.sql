CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE passengers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    name TEXT NOT NULL,
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL
);

CREATE TABLE drivers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    name TEXT NOT NULL,
    license_id TEXT NOT NULL,
    available BOOLEAN NOT NULL,
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL,
    car_make TEXT NOT NULL,
    car_model TEXT NOT NULL,
    car_year INT NOT NULL
);

CREATE TABLE rides (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    passenger_id UUID NOT NULL REFERENCES passengers(id),
    driver_id UUID REFERENCES drivers(id),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    status TEXT NOT NULL
);