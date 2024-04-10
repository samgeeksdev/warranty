CREATE TABLE manufacturers (
                               id SERIAL PRIMARY KEY,
                               name VARCHAR(255) NOT NULL UNIQUE,  -- Enforce unique manufacturer names
                               description TEXT DEFAULT NULL  -- Optional description of the manufacturer
);

CREATE TABLE warranty_types (
                                id SERIAL PRIMARY KEY,
                                name VARCHAR(255) NOT NULL UNIQUE
);
CREATE TABLE permissions (
                             id SERIAL PRIMARY KEY,
                             name VARCHAR(255) NOT NULL UNIQUE,
                             description TEXT,
                             created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE roles (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL UNIQUE,
                       description TEXT,
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE users (
                       id BIGSERIAL PRIMARY KEY,
                       username VARCHAR(255) NOT NULL UNIQUE,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       password_hash VARCHAR(255) NOT NULL,
                       first_name VARCHAR(255),
                       last_name VARCHAR(255),
                       phone_number VARCHAR(255),
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                       active BOOLEAN NOT NULL DEFAULT TRUE,
                       CONSTRAINT chk_email CHECK (email ~ '^[\\w!#$%&''*+/=?^`{|}~-]+(?:\\.[\\w!#$%&''*+/=?^`{|}~-]+)*@(?:[A-Z0-9-]+\\.)+[A-Z]{2,}$')
);

CREATE TABLE user_roles (
                            id BIGSERIAL PRIMARY KEY,
                            user_id BIGINT NOT NULL,
                            role_id INT NOT NULL,
                            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            FOREIGN KEY (user_id) REFERENCES users(id),
                            FOREIGN KEY (role_id) REFERENCES roles(id)
);

CREATE TABLE product_categories (
                                    id BIGSERIAL PRIMARY KEY,
                                    name VARCHAR(255) NOT NULL UNIQUE,
                                    description TEXT DEFAULT NULL,
                                    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
);
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
                          updated_at TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                          warranty_months INT,
                          retired BOOLEAN NOT NULL DEFAULT FALSE,
                          CONSTRAINT chk_price CHECK (price > 0) -- Ensures positive price
);
CREATE TABLE product_category_associations (
                                               product_id BIGINT NOT NULL,
                                               category_id BIGINT NOT NULL,
                                               is_primary BOOLEAN DEFAULT FALSE,
                                               FOREIGN KEY (product_id) REFERENCES products(id),
                                               FOREIGN KEY (category_id) REFERENCES product_categories(id),
                                               PRIMARY KEY (product_id, category_id)
);
CREATE TABLE customers (
                           id BIGSERIAL PRIMARY KEY,  -- Auto-incrementing primary key for customer identification
                           user_id BIGINT DEFAULT NULL,  -- Foreign key referencing users.id (customer can be linked to a user account)
                           name VARCHAR(255) NOT NULL,  -- Customer name (required)
                           email VARCHAR(255) NOT NULL UNIQUE, -- Customer email (required, unique for each customer)
                           phone_number VARCHAR(255) DEFAULT NULL, -- Optional phone number
                           address TEXT,  -- Optional customer address details
                           created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Timestamp of customer record creation (automatically set)
                           CONSTRAINT fk_customer_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL  -- Foreign key constraint with ON DELETE SET NULL for orphaned customer records
);

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
CREATE TABLE claim_statuses (
                                id SERIAL PRIMARY KEY,
                                name VARCHAR(255) NOT NULL UNIQUE
);
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
CREATE TABLE repair_representatives (
                                        id SERIAL PRIMARY KEY,
                                        user_id BIGINT REFERENCES users(id) ON DELETE SET NULL, -- Optional, links to user account (allows soft delete)
                                        name VARCHAR(255) NOT NULL,
                                        phone_number VARCHAR(255),
                                        email VARCHAR(255) UNIQUE, -- Optional, for contacting representatives
                                        active BOOLEAN NOT NULL DEFAULT TRUE, -- Tracks active/inactive status
                                        certification VARCHAR(255), -- Optional, for storing relevant certifications
                                        specialization VARCHAR(255), -- Optional, for specific repair expertise
                                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                        updated_at TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE availabilities (
                                id SERIAL PRIMARY KEY,
                                repair_representative_id BIGINT NOT NULL REFERENCES repair_representatives(id),
                                availability_type VARCHAR(255), -- e.g., "Full Time", "Part Time", "On-Call"
                                availability_details TEXT,  -- Optional details about availability schedule
                                created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                updated_at TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE skills (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL UNIQUE,
                        description TEXT,
                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE repair_representative_skills (
                                              repair_representative_id BIGINT NOT NULL REFERENCES repair_representatives(id),
                                              skill_id BIGINT NOT NULL REFERENCES skills(id),
                                              PRIMARY KEY (repair_representative_id, skill_id)
);
