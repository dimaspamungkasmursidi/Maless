CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    google_id VARCHAR(255) NOT NULL UNIQUE, 
    name VARCHAR(255) NOT NULL,             
    email VARCHAR(255) NOT NULL UNIQUE,     
    avatar_url TEXT,                        
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

