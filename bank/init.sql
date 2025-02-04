CREATE TABLE customers (
    customer_id  INT PRIMARY KEY AUTO_INCREMENT,
    name         VARCHAR(255) NOT NULL,
    date_of_birth DATE NOT NULL,
    city         VARCHAR(100) NOT NULL,
    zipcode      VARCHAR(20) NOT NULL,
    status       TINYINT NOT NULL
);

CREATE TABLE accounts (
    account_id    INT PRIMARY KEY AUTO_INCREMENT,
    customer_id   INT NOT NULL,
    opening_date  DATE NOT NULL,
    account_type  VARCHAR(50) NOT NULL,
    amount        DECIMAL(15,2) NOT NULL,
    status        TINYINT NOT NULL,
    
    FOREIGN KEY (customer_id) REFERENCES customer(customer_id) 
);


INSERT INTO customers (name, date_of_birth, city, zipcode, status) VALUES
('Alice Johnson', '1985-03-15', 'New York', '10001', 1),
('Bob Smith', '1990-07-22', 'Los Angeles', '90001', 1),
('Catherine Zeta', '1978-11-02', 'Chicago', '60601', 0),
('David Brown', '1992-12-25', 'Houston', '77001', 1),
('Eva Green', '1983-08-18', 'Phoenix', '85001', 1),
('Frank White', '1975-05-05', 'Philadelphia', '19101', 0),
('Grace Lee', '1999-02-28', 'San Antonio', '78201', 1),
('Henry Ford', '1988-09-09', 'San Diego', '92101', 1),
('Ivy Banks', '1995-01-15', 'Dallas', '75201', 0),
('Jack Black', '1980-06-06', 'San Jose', '95101', 1);

INSERT INTO Account (account_id, customer_id, opening_date, account_type, amount, status) VALUES
(1, 1, '2023-01-15', 'Savings', 5000.00, 1),
(2, 1, '2023-03-10', 'Checking', 12000.50, 1),
(3, 1, '2022-07-20', 'Savings', 3000.75, 0),
(4, 2, '2021-11-05', 'Fixed Deposit', 25000.00, 1),
(5, 2, '2024-02-01', 'Savings', 1000.00, 1),
(6, 3, '2023-05-18', 'Checking', 7500.25, 0),
(7, 3, '2022-09-30', 'Savings', 8900.00, 1),
(8, 4, '2021-12-22', 'Fixed Deposit', 50000.00, 1),
(9, 4, '2023-08-14', 'Checking', 2000.00, 1),
(10, 5, '2024-01-05', 'Savings', 6500.50, 0);


