CREATE TABLE IF not Exists users (
    id SERIAL PRIMARY KEY,                -- Auto-incremental ID
    username VARCHAR(255) NOT NULL UNIQUE,       -- Username, required field
    email VARCHAR(255) NOT NULL UNIQUE,   -- Email, must be unique
    password bytea NOT NULL,               -- Password, required field
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Auto-set to current time on creation
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Auto-set to current time on creation
);
