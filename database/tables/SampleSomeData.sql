INSERT INTO manufacturers (name, description)
VALUES ('Acme Corporation', 'A leading manufacturer of household appliances');

INSERT INTO manufacturers (name, description)
VALUES ('TechTronics Inc.', 'A global electronics company');


INSERT INTO warranty_types (name)
VALUES ('Standard Warranty');

INSERT INTO warranty_types (name)
VALUES ('Extended Warranty');

INSERT INTO permissions (name, description)
VALUES ('create_user', 'Permission to create new users');

INSERT INTO permissions (name, description)
VALUES ('edit_product', 'Permission to edit product information');

INSERT INTO roles (name, description)
VALUES ('admin', 'Full system access');

INSERT INTO roles (name, description)
VALUES ('customer_support', 'Ability to view and manage customer support tickets');

INSERT INTO users (username, email, password_hash, first_name, last_name, created_at)
VALUES ('johndoe', 'john.doe@example.com', 'hashed_password', 'John', 'Doe', CURRENT_TIMESTAMP);

INSERT INTO users (username, email, password_hash, first_name, last_name, created_at)
VALUES ('janesmith', 'jane.smith@example.com', 'hashed_password2', 'Jane', 'Smith', CURRENT_TIMESTAMP);

-- Assign admin role to the first user (johndoe)
INSERT INTO user_roles (user_id, role_id)
VALUES (1, 1);

INSERT INTO product_categories (name, description)
VALUES ('Laptops', 'A wide range of laptops for various needs');

INSERT INTO product_categories (name, description)
VALUES ('Smartphones', 'The latest smartphones with cutting-edge features');

INSERT INTO products (name, slug, description, category_id, manufacturer_id, model, sku, price, stock, warranty_months)
VALUES ('Acme Laptop 1000', 'acme-laptop-1000', 'A powerful laptop for everyday tasks', 1, 1, 'XL1000', 'ACME-LT-1000', 499.99, 10, 12);

INSERT INTO products (name, slug, description, category_id, manufacturer_id, model, sku, price, stock, warranty_months)
VALUES ('TechTronics Phone X', 'tectronics-phone-x', 'The latest smartphone with a groundbreaking camera', 2, 2, 'TX7', 'TEC-PH-X7', 899.99, 20, 24);

-- Associate the first product (Acme Laptop 1000) with the Laptops category (set as primary)
INSERT INTO product_category_associations (product_id, category_id, is_primary)
VALUES (1, 1, TRUE);

INSERT INTO customers (name, email, phone_number)
VALUES ('John Doe', 'johndoe@customeremail.com', '123-456-7890');

INSERT INTO customers (name, email)
VALUES ('Jane Smith', 'janesmith@customeremail.com');

-- Assuming John Doe (user ID 1) registered the warranty
INSERT INTO warranties (product_id, warranty_type_id, duration_months, start_date, registered_by, customer_id)
VALUES (1, 1, 12, CURRENT_TIMESTAMP - interval '1 month', 1, 1);

INSERT INTO claim_statuses (name)
VALUES ('Open');

INSERT INTO claim_statuses (name)
VALUES ('Pending Approval');

INSERT INTO claim_statuses (name)
VALUES ('Approved');

INSERT INTO claim_statuses (name)
VALUES ('Denied');

INSERT INTO claim_statuses (name)
VALUES ('Resolved');


INSERT INTO claims (warranty_id, description, claim_date, status_id)
VALUES (1, 'Faulty display', CURRENT_TIMESTAMP - interval '2 weeks', 2);  -- Pending Approval

INSERT INTO claim_audits (claim_id, modified_field, old_value, new_value, modified_by, modified_at)
VALUES (1, 'status_id', '{"id": 1}', '{"id": 2}', 1, CURRENT_TIMESTAMP);

INSERT INTO repair_representatives (name, email, phone_number, active)
VALUES ('Alice Technician', 'alice.technician@repair.com', '555-123-4567', TRUE);

INSERT INTO repair_representatives (name, email, phone_number, active)
VALUES ('Bob Fixit', 'bob.fixit@repair.com', '555-789-0123', TRUE);

INSERT INTO availabilities (repair_representative_id, availability_type)
VALUES (1, 'Full Time');

INSERT INTO availabilities (repair_representative_id, availability_type)
VALUES (2, 'Part Time');



INSERT INTO skills (name, description)
VALUES ('Laptop Repair', 'Experience in diagnosing and repairing laptop hardware issues');

INSERT INTO skills (name, description)
VALUES ('Smartphone Repair', 'Expertise in fixing smartphone screens, batteries, and other components');



