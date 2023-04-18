CREATE TABLE film_categories (
    film_id INT,
    category_id INT,
    PRIMARY KEY (film_id, category_id),
    FOREIGN KEY (film_id) REFERENCES films (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE RESTRICT ON UPDATE CASCADE
);