CREATE DATABASE IF NOT EXISTS db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE db;

-- Users Table
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Books Table
CREATE TABLE IF NOT EXISTS books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    author VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Orders Tables
CREATE TABLE IF NOT EXISTS orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    book_id INT NOT NULL,
    quantity INT NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(id),
    FOREIGN KEY (book_id) REFERENCES Books(id)
);

INSERT INTO users (username, email, password) VALUES
('Alice', 'alice@example.com', 'hashed_password1'),
('Bob', 'bob@example.com', 'hashed_password2'),
('Charlie', 'charlie@example.com', 'hashed_password3');

INSERT INTO books (title, author, price, stock) VALUES
('Go Programming', 'John Doe', 49.99, 100),
('Mastering GORM', 'Jane Smith', 59.99, 50),
('Database Design', 'Michael Brown', 39.99, 75);

INSERT INTO orders (user_id, book_id, quantity, total_price) VALUES
(1, 1, 2, 99.98),   -- Alice购买了2本Go Programming，总价99.98
(2, 2, 1, 59.99),   -- Bob购买了1本Mastering GORM，总价59.99
(3, 3, 3, 119.97);  -- Charlie购买了3本Database Design，总价119.97
