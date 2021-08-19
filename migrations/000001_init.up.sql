CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    birth_date TIMESTAMP NOT NULL,
    gender VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    address VARCHAR(200)
);

-- FASTEST AND WORST WAY TO FILL DB
INSERT INTO customers(first_name, last_name, birth_date, gender, email)
VALUES ('TestFirstNameOne', 'TestLastNameOne', TIMESTAMP '2001-10-22 10:23:54.000000', 'Male', 'testOne@mail.ru', 'addressOne');

INSERT INTO customers(first_name, last_name, birth_date, gender, email)
VALUES ('TestFirstNameTwo', 'TestLastNameTwo', TIMESTAMP '2001-10-23 10:23:54.000000', 'Female', 'testTwo@mail.ru', 'addressTwo');