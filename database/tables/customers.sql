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
