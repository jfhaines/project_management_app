USE project_management_app;
DROP TABLE IF EXISTS card_checklist_items;

CREATE TABLE card_checklist_items (
  id INTEGER AUTO_INCREMENT,
  text VARCHAR(20) NOT NULL,
  position INTEGER NOT NULL,
  card_id INTEGER NOT NULL,
  FOREIGN KEY (card_id) REFERENCES cards(id),
  PRIMARY KEY (id)
);

INSERT INTO card_checklist_items (text, position, card_id)
VALUES
  ('Checklist item 1', 1, 1),
  ('Checklist item 2', 2, 1),
  ('Checklist item 3', 3, 1),
  ('Checklist item 1', 1, 2),
  ('Checklist item 2', 2, 2),
  ('Checklist item 3', 3, 2),
  ('Checklist item 1', 1, 3),
  ('Checklist item 2', 2, 3),
  ('Checklist item 3', 3, 3);