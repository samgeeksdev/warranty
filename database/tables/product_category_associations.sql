CREATE TABLE product_category_associations (
                                               product_id BIGINT NOT NULL,
                                               category_id BIGINT NOT NULL,
                                               is_primary BOOLEAN DEFAULT FALSE,  -- Flag for primary category (optional)
                                               FOREIGN KEY (product_id) REFERENCES products(id),
                                               FOREIGN KEY (category_id) REFERENCES product_categories(id),
                                               PRIMARY KEY (product_id, category_id)  -- Composite primary key for many-to-many relationship
);