-- +goose Up
CREATE TABLE flavors (
    id UUID PRIMARY KEY, 
    name VARCHAR(255) NOT NULL,
    add_price NUMERIC(8,2) NOT NULL
);
    
-- +goose Down
DROP TABLE flavors;
