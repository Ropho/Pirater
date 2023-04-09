START TRANSACTION;

CREATE TABLE user_rights (
    right_id INT REFERENCES rights (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    user_id INT REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (right_id, user_id)
);

INSERT INTO user_rights (right_id, user_id)
VALUES (4, 1);

COMMIT;