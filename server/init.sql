CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    published_date DATE,
    genre VARCHAR(100),
    price NUMERIC(10, 2) NOT NULL,
    stock INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO books (title, author, published_date, genre, price, stock)
VALUES 
('To Kill a Mockingbird', 'Harper Lee', '1960-07-11', 'Fiction', 9.99, 50),
('1984', 'George Orwell', '1949-06-08', 'Dystopian', 14.99, 30),
('Pride and Prejudice', 'Jane Austen', '1813-01-28', 'Romance', 12.99, 20),
('The Great Gatsby', 'F. Scott Fitzgerald', '1925-04-10', 'Fiction', 10.99, 40),
('Moby-Dick', 'Herman Melville', '1851-10-18', 'Adventure', 11.49, 15),
('The Catcher in the Rye', 'J.D. Salinger', '1951-07-16', 'Fiction', 8.99, 25),
('The Hobbit', 'J.R.R. Tolkien', '1937-09-21', 'Fantasy', 13.49, 35),
('War and Peace', 'Leo Tolstoy', '1869-01-01', 'Historical', 19.99, 10),
('The Alchemist', 'Paulo Coelho', '1988-04-15', 'Adventure', 16.99, 45),
('Crime and Punishment', 'Fyodor Dostoevsky', '1866-01-01', 'Mystery', 14.49, 20);
