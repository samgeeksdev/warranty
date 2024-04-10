CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL UNIQUE,
                          slug VARCHAR(255) UNIQUE,
                          description TEXT,
                          category_id INT NOT NULL REFERENCES product_categories(id),
                          manufacturer_id INT REFERENCES manufacturers(id),
                          model VARCHAR(255),
                          sku VARCHAR(255) NOT NULL UNIQUE,
                          price NUMERIC(10,2) NOT NULL,
                          stock INT NOT NULL DEFAULT 0,
                          created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT NULL,  -- Removed default for update tracking
                          warranty_months INT,
                          retired BOOLEAN NOT NULL DEFAULT FALSE,
                          CONSTRAINT chk_price CHECK (price > 0) -- Ensures positive price
);
