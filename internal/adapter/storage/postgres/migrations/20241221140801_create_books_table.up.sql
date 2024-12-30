CREATE TABLE IF NOT EXISTS books (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    author TEXT NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    cover TEXT
);