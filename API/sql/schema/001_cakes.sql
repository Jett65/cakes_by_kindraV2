-- +goose Up
CREATE TABLE cakes (
    id UUID PRIMARY KEY,
    type VARCHAR(50) NOT NULL,
    layer_number SMALLINT NOT NULL,
    tiere_number SMALLINT NOT NULL,
    size VARCHAR(50) NOT NULL,
    price NUMERIC(8,2) NOT NULL
);

-- +goose Down
DROP TABLE cakes;

