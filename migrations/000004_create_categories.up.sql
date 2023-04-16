BEGIN;

CREATE TABLE categories (
    id INT PRIMARY KEY AUTO_INCREMENT,
    category VARCHAR (40) NOT NULL UNIQUE
);

INSERT INTO categories (category) VALUES
    ("drama"),
    ("melodrama"),
    ("horror"),
    ("thriller"),
    ("adventure"),
    ("crime"),
    ("historical"),
    ("comedy"),
    ("scientific"),
    ("biographical"),
    ("action movie"),
    ("music"),
    ("musical"),
    ("fantasy"),
    ("sport")
    ;


COMMIT;