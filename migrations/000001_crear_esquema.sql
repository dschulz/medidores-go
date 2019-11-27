

-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS medidores
(
    id             SERIAL NOT NULL PRIMARY KEY,
    numero         TEXT
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS medidores;