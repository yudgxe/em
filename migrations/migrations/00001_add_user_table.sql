-- +goose Up
-- +goose StatementBegin
CREATE TYPE gender_type AS ENUM('male', 'female', 'other');

CREATE TABLE users (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL CHECK (name <> ''),
    surname     VARCHAR(255) NOT NULL CHECK (surname <> ''),
    patronymic  VARCHAR(255) DEFAULT NULL CHECK (patronymic <> ''),
    age         INTEGER DEFAULT NULL CHECK (age > 0),
    nationality TEXT DEFAULT NULL CHECK (nationality <> ''),
    gender      gender_type DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TYPE gender_type;
-- +goose StatementEnd
