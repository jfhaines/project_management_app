USE project_management_app;
DROP TABLE IF EXISTS card_comments;

CREATE TABLE card_comments (
  id VARCHAR(40),
  text VARCHAR(1000) NOT NULL,
  datetime TIMESTAMP NOT NULL,
  user_id VARCHAR(40) NOT NULL,
  card_id VARCHAR(40) NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (card_id) REFERENCES cards(id),
  PRIMARY KEY (id)
);

INSERT INTO card_comments (text, time, user_id, card_id)
VALUES
  ('This is a comment', '2022-05-12 18:00:00', 1, 1),
  ('This is a comment', '2022-05-12 18:00:00', 1, 2),
  ('This is a comment', '2022-05-12 18:00:00', 1, 3);