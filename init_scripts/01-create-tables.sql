-- create users table
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- create petitions table
CREATE TABLE IF NOT EXISTS petitions (
    id VARCHAR(255) PRIMARY KEY,
		user_id VARCHAR(255) REFERENCES user(id) ON DELETE CASCADE,
		game_name VARCHAR(255) NOT NULL,
		game_url TEXT NOT NULL,
		status VARCHAR(50) DEFAULT 'pending', -- pending, testing, completed
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
