BEGIN;

DO $$ BEGIN
  CREATE TYPE status_types AS ENUM('wait', 'onprogress', 'done');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

ALTER TABLE meta_orders ADD COLUMN status status_types DEFAULT 'wait' NOT NULL;
ALTER TABLE orders ADD COLUMN status status_types DEFAULT 'wait' NOT NULL;

COMMIT;
