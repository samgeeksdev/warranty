CREATE TABLE claim_audits (
                              id SERIAL PRIMARY KEY,
                              claim_id INT NOT NULL,
                              modified_field VARCHAR(255) NOT NULL,
                              old_value JSON DEFAULT NULL,
                              new_value JSON DEFAULT NULL,
                              modified_by INT NOT NULL,
                              modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              FOREIGN KEY (claim_id) REFERENCES claims(id)
);
