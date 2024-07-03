-- +migrate Up
CREATE TABLE trips (
    id uuid NOT NULL DEFAULT uuid_generate_v4 (),
    name varchar(255),
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL,
    CONSTRAINT trips_pkey PRIMARY KEY (id),
    user_id uuid NOT NULL
);

CREATE INDEX idx_tripss_fk_user_uid ON trips USING btree (user_id);

ALTER TABLE trips
    ADD CONSTRAINT trips_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE;

-- +migrate Down
DROP TABLE IF EXISTS trips;

