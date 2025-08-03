CREATE  TABLE  IF NOT EXISTS sets(
    id SERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL,
    name VARCHAR(64) NOT NULL,
    description TEXT
)