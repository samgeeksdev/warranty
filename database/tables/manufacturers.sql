CREATE TABLE manufacturers (
                               id SERIAL PRIMARY KEY,
                               name VARCHAR(255) NOT NULL UNIQUE,  -- Enforce unique manufacturer names
                               description TEXT DEFAULT NULL  -- Optional description of the manufacturer
);