USE project_management_app;
DROP TABLE IF EXISTS cards;

CREATE TABLE cards (
  id VARCHAR(40),
  title VARCHAR(200) NOT NULL,
  description VARCHAR(1500),
  due_at TIMESTAMP,
  column_id VARCHAR(40) NOT NULL,
  FOREIGN KEY (column_id) REFERENCES columns(id),
  PRIMARY KEY (id)
);

INSERT INTO cards (title, description, date_due, column_id)
VALUES
  ('Card 1', 'Description for card 1', '2022-05-12 18:00:00', 1),
  ('Card 2', 'Description for card 2', '2022-05-12 18:00:00', 2),
  ('Card 3', 'Description for card 3', '2022-05-12 18:00:00', 3);