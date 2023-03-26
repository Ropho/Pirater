CREATE TABLE user_bookmarks (
    film_id INT REFERENCES films (id) ON DELETE CASCADE ON UPDATE CASCADE,
    user_id INT REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (film_id, user_id)
);