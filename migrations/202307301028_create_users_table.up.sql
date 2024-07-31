CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    tags TEXT[],
    bio TEXT,
    social TEXT[],
    image TEXT,
    password TEXT NOT NULL,
);