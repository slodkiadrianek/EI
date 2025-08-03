CREATE  TABLE IF NOT EXISTS elements (
    id SERIAL PRIMARY KEY,
    set_id INTEGER NOT NULL,
    english VARCHAR(64) NOT NULL,
    polish VARCHAR(64) NOT NULL,
    example_sentence TEXT NOT NULL,
    synonym VARCHAR(64), NOT NULL
);