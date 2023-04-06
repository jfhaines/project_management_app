USE project_management_app;
DROP TABLE IF EXISTS columns;

CREATE TABLE columns (
  id INTEGER AUTO_INCREMENT,
  name VARCHAR(60) NOT NULL,
  board_id INTEGER NOT NULL,
  FOREIGN KEY (board_id) REFERENCES boards(id),
  PRIMARY KEY (id)
);

INSERT INTO columns (name, board_id)
VALUES
  ('To Do', 1),
  ('In Progress', 1),
  ('Review', 1),
  ('Completed', 1);