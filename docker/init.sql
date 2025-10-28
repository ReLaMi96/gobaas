-- Initial database setup for gobaas
-- This file will be executed when the PostgreSQL container starts for the first time

-- Create any additional extensions you might need
-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Set timezone
SET timezone = 'Europe/Budapest';

-- You can add any initial data or additional setup here
-- For example:
-- INSERT INTO your_table (column1, column2) VALUES ('value1', 'value2');

COMMENT ON DATABASE gobaas IS 'GoBaas application database';