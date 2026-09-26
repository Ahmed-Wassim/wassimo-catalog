-- +goose Up
CREATE TABLE IF NOT EXISTS restaurants (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT        NOT NULL,
    status      TEXT        NOT NULL DEFAULT 'active',  -- e.g. active, inactive, pending
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS restaurants;