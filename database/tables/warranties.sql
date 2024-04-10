CREATE TABLE warranties (
                            id BIGSERIAL PRIMARY KEY,
                            product_id BIGINT NOT NULL,
                            warranty_type_id INT NOT NULL,
                            duration_months INT,
                            start_date TIMESTAMP NOT NULL,
                            end_date TIMESTAMP DEFAULT NULL,
                            registered_by BIGINT NOT NULL,
                            customer_id BIGINT DEFAULT NULL,
                            FOREIGN KEY (product_id) REFERENCES products(id),
                            FOREIGN KEY (registered_by) REFERENCES users(id),
                            FOREIGN KEY (warranty_type_id) REFERENCES warranty_types(id),
                            FOREIGN KEY (customer_id) REFERENCES customers(id)
);