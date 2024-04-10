CREATE TABLE warranty_audits (
                                 id SERIAL PRIMARY KEY,
                                 warranty_id BIGINT NOT NULL,
                                 modified_field VARCHAR(255) NOT NULL,
                                 old_value JSON DEFAULT NULL,
                                 new_value JSON DEFAULT NULL,
                                 modified_by BIGINT NOT NULL,
                                 modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 FOREIGN KEY (warranty_id) REFERENCES warranties(id)
);