CREATE TABLE claims (
                        id SERIAL PRIMARY KEY,
                        warranty_id BIGINT NOT NULL,
                        description TEXT NOT NULL,
                        claim_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        status_id INT NOT NULL,
                        resolution TEXT,
                        claim_files_json TEXT,
                        FOREIGN KEY (warranty_id) REFERENCES warranties(id),
                        FOREIGN KEY (status_id) REFERENCES claim_statuses(id)
);