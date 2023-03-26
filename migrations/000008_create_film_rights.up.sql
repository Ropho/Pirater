CREATE TABLE film_rights (
    right_id INT REFERENCES rights (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    film_id INT REFERENCES films (id) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (right_id, film_id)
);