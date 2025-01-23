
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION generate_uuid_v7() RETURNS uuid AS $$
DECLARE
    ts BIGINT;
    rand_bytes BYTEA;
    uuid_v7 UUID;
BEGIN
    -- Get the current Unix timestamp in milliseconds
    ts := EXTRACT(EPOCH FROM clock_timestamp()) * 1000;

    -- Generate 10 random bytes
    rand_bytes := gen_random_bytes(10);

    -- Construct the UUID v7
    uuid_v7 := (
        lpad(to_hex(ts), 12, '0') || -- 48 bits of timestamp
        '7' || -- 4 bits of version (UUIDv7)
        substring(to_hex(get_byte(rand_bytes, 0)), 2, 1) || -- 4 bits of random data
        substring(to_hex(rand_bytes), 2) -- 80 bits of random data
    )::uuid;

    RETURN uuid_v7;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE passengers (
    id UUID PRIMARY KEY DEFAULT generate_uuid_v7(),
    name TEXT NOT NULL,
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL
);

CREATE TABLE drivers (
    id UUID PRIMARY KEY DEFAULT generate_uuid_v7(),
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
    id UUID PRIMARY KEY DEFAULT generate_uuid_v7(),
    passenger_id UUID NOT NULL REFERENCES passengers(id),
    driver_id UUID REFERENCES drivers(id),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    status TEXT NOT NULL
);