CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR (30) UNIQUE NOT NULL,
    pass VARCHAR (75) NOT NULL
);