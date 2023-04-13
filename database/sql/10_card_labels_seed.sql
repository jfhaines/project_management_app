USE project_management_app;
DROP TABLE IF EXISTS card_labels;

CREATE TABLE card_labels (
  id VARCHAR(40),
  text VARCHAR(20) NOT NULL,
  colour VARCHAR(30) NOT NULL,
  card_id VARCHAR(40) NOT NULL,
  FOREIGN KEY (card_id) REFERENCES cards(id),
  PRIMARY KEY (id)
);

INSERT INTO card_labels (text, colour, card_id)
VALUES
  ('Label 1', 'Red', 1),
  ('Label 2', 'Blue', 1),
  ('Label 1', 'Red', 2),
  ('Label 2', 'Blue', 2),
  ('Label 1', 'Red', 3),
  ('Label 2', 'Blue', 3);