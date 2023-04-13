USE project_management_app;
DROP TABLE IF EXISTS card_users;

CREATE TABLE card_users (
  id VARCHAR(40),
  user_id VARCHAR(40) NOT NULL,
  card_id VARCHAR(40) NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (card_id) REFERENCES cards(id),
  PRIMARY KEY (id)
);

INSERT INTO card_users (user_id, card_id)
VALUES
  (1, 1),
  (2, 2),
  (3, 3),
  (2, 1);
