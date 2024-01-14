-- +goose Up
Create Table filling (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price NUMERIC(8,2) NOT NULL
);

-- +goose Down
DROP TABLE filling;
