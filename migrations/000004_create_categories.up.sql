BEGIN;

CREATE TABLE categories (
    id INT PRIMARY KEY AUTO_INCREMENT,
    category VARCHAR (20) NOT NULL UNIQUE
);

INSERT INTO categories (category) VALUES
    ("DRAMA"),
    ("HORROR"),
    ("THRILLER"),
    ("ADVENTURE")
    ;


COMMIT;