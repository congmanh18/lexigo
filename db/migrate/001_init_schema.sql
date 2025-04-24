-- +goose Up
CREATE TABLE china_entries (
    id SERIAL PRIMARY KEY,
    traditional TEXT,
    simplified TEXT,
    pinyin TEXT,
    definition TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- +goose Down
DROP TABLE china_entries;
