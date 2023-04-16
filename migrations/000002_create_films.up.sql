CREATE TABLE films (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR (300) UNIQUE NOT NULL,
    hash INT UNSIGNED UNIQUE NOT NULL,
    description TEXT NOT NULL,
    afisha_url VARCHAR (300) NOT NULL,
    header_url VARCHAR (300) NOT NULL,
    video_url VARCHAR (300) NOT NULL
    /* cadre_url VARCHAR (300) NOT NULL */
);