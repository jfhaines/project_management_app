USE project_management_app;
DROP TABLE IF EXISTS boards;

CREATE TABLE boards (
  id INTEGER AUTO_INCREMENT,
  name VARCHAR(120) NOT NULL,
  background VARCHAR(255),
  PRIMARY KEY (id)
);

INSERT INTO boards (name)
VALUES 
  ('Board 1'),
  ('Board 2'),
  ('Board 3');