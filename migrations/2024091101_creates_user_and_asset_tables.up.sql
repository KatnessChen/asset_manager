-- +migrate Up
CREATE TABLE users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE asset_types (
    asset_type_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT
);

CREATE TABLE assets (
    asset_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    asset_type_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (asset_type_id) REFERENCES asset_types(asset_type_id) ON DELETE RESTRICT
);

CREATE TABLE asset_records (
    record_id INT AUTO_INCREMENT PRIMARY KEY,
    asset_id INT NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    record_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (asset_id) REFERENCES assets(asset_id) ON DELETE CASCADE
);

-- Create an index on the record_date for faster queries
CREATE INDEX idx_asset_records_record_date ON asset_records(record_date);

-- +migrate Down
-- DROP TABLE IF EXISTS asset_records;
-- DROP TABLE IF EXISTS assets;
-- DROP TABLE IF EXISTS asset_types;
-- DROP TABLE IF EXISTS users;
