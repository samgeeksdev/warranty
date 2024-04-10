CREATE TABLE availabilities (
                                id SERIAL PRIMARY KEY,
                                repair_representative_id BIGINT NOT NULL REFERENCES repair_representatives(id),
                                availability_type VARCHAR(255),
                                availability_details TEXT,
                                created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                updated_at TIMESTAMP DEFAULT NULL
);
