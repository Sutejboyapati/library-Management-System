-- Run this in your MySQL to create the database and tables.
-- Use your own database name if you prefer (then set DB_NAME in .env).

CREATE DATABASE IF NOT EXISTS library_db;
USE library_db;

CREATE TABLE IF NOT EXISTS Users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  role VARCHAR(50) NOT NULL DEFAULT 'user',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Books (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(500) NOT NULL,
  author VARCHAR(255) NOT NULL,
  isbn VARCHAR(50),
  genre VARCHAR(100),
  language VARCHAR(100),
  shelf_number VARCHAR(50),
  available_copies INT NOT NULL DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS borrowingrecords (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  book_id INT NOT NULL,
  borrowed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  due_date TIMESTAMP NOT NULL,
  returned_at TIMESTAMP NULL,
  FOREIGN KEY (user_id) REFERENCES Users(id),
  FOREIGN KEY (book_id) REFERENCES Books(id)
);

-- Optional: add an admin user (password: admin123)
-- INSERT INTO Users (username, password, role) VALUES ('admin', '$2a$10$...', 'admin');
-- Generate bcrypt hash for 'admin123' and replace above.
