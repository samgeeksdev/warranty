CREATE TABLE repair_representatives (
                                        id SERIAL PRIMARY KEY,
                                        user_id BIGINT REFERENCES users(id) ON DELETE SET NULL,
                                        name VARCHAR(255) NOT NULL,
                                        phone_number VARCHAR(255),
                                        email VARCHAR(255) UNIQUE,
                                        active BOOLEAN NOT NULL DEFAULT TRUE,
                                        certification VARCHAR(255),
                                        specialization VARCHAR(255),
                                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                        updated_at TIMESTAMP DEFAULT NULL
);
