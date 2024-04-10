CREATE TABLE permissions (
                             id SERIAL PRIMARY KEY,  -- Use SERIAL for auto-incrementing primary key in PostgreSQL
                             name VARCHAR(255) NOT NULL UNIQUE,
                             description TEXT,
                             created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
