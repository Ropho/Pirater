START TRANSACTION;

CREATE TABLE user_rights (
    right_id INT,
    user_id INT,
    PRIMARY KEY (right_id, user_id),
    FOREIGN KEY (right_id) REFERENCES rights (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
);

INSERT INTO user_rights (right_id, user_id)
VALUES (4, 1);

COMMIT;