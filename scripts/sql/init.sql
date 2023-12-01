-- This is the initial script to create the database required for the egghead service to use

-- Create the user with the required access
-- CREATE USER IF NOT EXISTS admin WITH PASSWORD 'admin';
DO
$do$
BEGIN
   IF EXISTS (
      SELECT FROM pg_catalog.pg_roles
      WHERE  rolname = 'admin') THEN

      RAISE NOTICE 'Role "admin" already exists. Skipping.';
   ELSE
      BEGIN   -- nested block
         CREATE ROLE admin LOGIN PASSWORD 'admin';
      EXCEPTION
         WHEN duplicate_object THEN
            RAISE NOTICE 'Role "admin" was just created by a concurrent transaction. Skipping.';
      END;
   END IF;
END
$do$;

-- Create the database
-- CREATE DATABASE IF NOT EXISTS egghead;
-- CREATE OR REPLACE FUNCTION create_database_if_not_exists()
-- RETURNS void 
-- LANGUAGE plpgsql
-- AS $$
-- BEGIN
--     IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'egghead') THEN
--         CREATE DATABASE egghead;
--     END IF;
-- END $$;

-- -- Call the function
-- SELECT create_database_if_not_exists();
CREATE DATABASE egghead WITH OWNER admin;

-- BEGIN
--    IF NOT EXISTS (
--       SELECT 1 FROM pg_database WHERE datname = 'egghead'
--    ) THEN
--       CREATE DATABASE egghead;
--    END IF;
-- END


-- Grant all access to the user
-- GRANT ALL PRIVILEGES ON DATABASE egghead TO admin;

-- Switch to the newly created database
-- \c egghead;

-- Create products table
-- CREATE TABLE IF NOT EXISTS products (
--     id SERIAL PRIMARY KEY,
--     uid VARCHAR(50) UNIQUE,
--     slug VARCHAR(50) UNIQUE NOT NULL,
--     name VARCHAR(50) NOT NULL
-- );

-- Create user tables
-- CREATE TABLE IF NOT EXISTS users (
--     id SERIAL PRIMARY KEY,
--     `username` VARCHAR(255) NOT NULL,
--     name VARCHAR(255) NOT NULL,
--     product_id INT NOT NULL,
--     archived BOOLEAN DEFAULT false,
--     balance INT DEFAULT 0,
--     external_user_id VARCHAR(255) NOT NULL,
--    --  created_on TIMESTAMP DEFAULT current_timestamp(),
--     FOREIGN KEY (product_id) REFERENCES products (id)
-- );

-- Index for id column
-- CREATE UNIQUE INDEX IF NOT EXISTS idx_id ON users(id);

-- Index for username column
-- CREATE UNIQUE INDEX IF NOT EXISTS idx_username ON users(username);
-- Index for the combination of product_id and username
-- CREATE UNIQUE INDEX IF NOT EXISTS idx_product_username ON users(product_id, username);
-- CREATE UNIQUE INDEX IF NOT EXISTS idx_product_id ON users(product_id, external_user_id);


-- Create coin_balance table
-- CREATE TABLE IF NOT EXISTS coin_balance (
--     id SERIAL PRIMARY KEY,
--     balance INT DEFAULT 0,
--     product_id INT NOT NULL,
--     user_id INT UNIQUE,
--     -- created_on TIMESTAMP NOT NULL,
--     -- updated_on TIMESTAMP NOT NULL,
--     FOREIGN KEY (product_id) REFERENCES products (id),
--     FOREIGN KEY (user_id) REFERENCES users (id)
-- );

-- Create index on coin_balance table
-- CREATE INDEX IF NOT EXISTS idx_product_users ON coin_balance (user_id, product_id);

-- Create transaction_history table
-- CREATE TABLE IF NOT EXISTS transaction_history (
-- 	id SERIAL PRIMARY KEY,
-- 	transaction_type VARCHAR(6),
-- 	product_id INT NOT NULL,
-- 	user_id INT UNIQUE NOT NULL,
-- 	amount INT NOT NULL,
--    reason: VARCHAR(255) NOT NULL,
-- 	timestamp TIMESTAMP DEFAULT current_timestamp,
-- 	FOREIGN KEY (product_id) REFERENCES products (id),
--    FOREIGN KEY (user_id) REFERENCES users (id)
-- );

-- Create index on transaction_history table
-- CREATE INDEX IF NOT EXISTS idx_transaction_history_product_user 
--     ON transaction_history (product_id, user_id);

