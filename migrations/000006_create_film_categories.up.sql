CREATE TABLE film_categories (
    film_id INT REFERENCES films (id) ON DELETE CASCADE ON UPDATE CASCADE,
    category_id INT REFERENCES categories (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    PRIMARY KEY (film_id, category_id)
);