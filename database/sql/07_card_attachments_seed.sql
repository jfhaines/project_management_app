USE project_management_app;
DROP TABLE IF EXISTS card_attachments;

CREATE TABLE card_attachments (
  id INTEGER AUTO_INCREMENT,
  filename VARCHAR(150) NOT NULL,
  time DATETIME NOT NULL,
  card_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  FOREIGN KEY (card_id) REFERENCES cards(id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  PRIMARY KEY (id)
);

INSERT INTO card_attachments (filename, time, card_id, user_id)
VALUES
  ('1', '2022-05-12 18:00:00', 1, 1),
  ('2', '2022-05-12 18:00:00', 2, 1),
  ('3', '2022-05-12 18:00:00', 2, 2);