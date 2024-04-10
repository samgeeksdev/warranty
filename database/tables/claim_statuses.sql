CREATE TABLE claim_statuses (
                                id SERIAL PRIMARY KEY,
                                name VARCHAR(255) NOT NULL UNIQUE,
                                description TEXT DEFAULT NULL,
                                sort_order INT DEFAULT NULL,
                                created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                updated_at TIMESTAMP DEFAULT NULL
);
