-- +goose Up
Create Table frosting (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    add_price NUMERIC(8,2) NOT NULL
);

-- +goose Down
DROP TABLE frosting;
