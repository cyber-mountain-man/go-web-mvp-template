-- init.sql (PostgreSQL)
-- Initialize database schema, roles, and tables

-- ✅ Create users table if not exists
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE
);

-- ✅ Add sample data (optional)
INSERT INTO users (name, email)
VALUES 
    ('Admin', 'admin@example.com'),
    ('Sample User', 'sample@example.com')
ON CONFLICT DO NOTHING;

-- ✅ Additional templates (uncomment and customize as needed)
-- CREATE TABLE IF NOT EXISTS posts (
--     id SERIAL PRIMARY KEY,
--     title TEXT NOT NULL,
--     body TEXT NOT NULL,
--     user_id INTEGER REFERENCES users(id)
-- );
