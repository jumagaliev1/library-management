CREATE TABLE IF NOT EXISTS users
(
    id serial PRIMARY KEY,
    first_name VARCHAR(128) NOT NULL,
    last_name VARCHAR(128) NOT NULL,
    email VARCHAR(256) UNIQUE NOT NULL,
    password VARCHAR(64) NOT NULL
);