CREATE TABLE IF NOT EXISTS recipes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    title TEXT,
    ingredient TEXT[],
    time INTEGER,
    process TEXT[] NOT NULL,
    tags TEXT[] NOT NULL,
    description TEXT,
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
    user_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);