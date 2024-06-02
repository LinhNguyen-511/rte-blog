CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS posts (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    uuid UUID NOT NULL DEFAULT uuid_generate_v4(), -- Generate UUID automatically,
    title VARCHAR(255) NOT NULL,
    author_id INT REFERENCES authors,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    published_at TIMESTAMP DEFAULT NULL
);