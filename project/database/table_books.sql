CREATE TABLE IF NOT EXISTS book (
    id SERIAL PRIMARY KEY,
    title TEXT,
    author TEXT,
    published_year INT
    isbn TEXT UNIQUE
);