BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS packages (
  id UUID DEFAULT uuid_generate_v1() PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  price BIGINT NOT NULL,
  quota BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS partners (
  id uuid DEFAULT uuid_generate_v1() PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,
  quota BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS meta_orders (
  id UUID DEFAULT uuid_generate_v1() PRIMARY KEY NOT NULL,
  partner_id UUID NOT NULL REFERENCES partners(id) ON DELETE CASCADE,
  start_at TIMESTAMP WITH TIME ZONE NOT NULL,
  description TEXT
);

DO $$ BEGIN
  CREATE TYPE transport_types AS ENUM('bike', 'taxi', 'food', 'package', 'truck');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS orders (
  id uuid DEFAULT uuid_generate_v1() PRIMARY KEY NOT NULL,
  meta_order_id UUID NOT NULL REFERENCES meta_orders(id) ON DELETE CASCADE,
  transport_type transport_types DEFAULT 'bike' NOT NULL,
  start_at TIMESTAMP WITH TIME ZONE NOT NULL,
  from_lang FLOAT NOT NULL,
  from_lat FLOAT NOT NULL,
  to_lang FLOAT NOT NULL,
  to_lat FLOAT NOT NULL,
  client_name VARCHAR(255),
  contact VARCHAR(255)
);

COMMIT;
