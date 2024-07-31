CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    bg_color VARCHAR(7) NOT NULL,
    txt_color VARCHAR(7) NOT NULL,
    image TEXT NOT NULL
);