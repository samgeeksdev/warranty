CREATE TABLE product_categories (
                                    id BIGSERIAL PRIMARY KEY,
                                    name VARCHAR(255) NOT NULL UNIQUE,  -- Enforce unique category names
                                    description TEXT DEFAULT NULL,  -- Optional description of the category
                                    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- Track creation time
                                    updated_at TIMESTAMP DEFAULT NULL  -- Optional update tracking
);
