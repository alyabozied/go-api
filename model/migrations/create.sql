CREATE TABLE users (
    id SERIAL PRIMARY KEY,                -- Auto-incremental ID
    username VARCHAR(255) NOT NULL,       -- Username, required field
    email VARCHAR(255) NOT NULL UNIQUE,   -- Email, must be unique
    password TEXT NOT NULL,               -- Password, required field
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Auto-set to current time on creation
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Auto-set to current time on creation
);
Insert Into users(username,email,password) values("abozied","abozied@gmail.com" , "hello")