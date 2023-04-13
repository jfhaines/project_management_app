USE project_management_app;
DROP TABLE IF EXISTS boards;

CREATE TABLE boards (
  id VARCHAR(40),
  name VARCHAR(120) NOT NULL,
  background VARCHAR(255),
  PRIMARY KEY (id)
);

INSERT INTO boards (name)
VALUES 
  ('Board 1'),
  ('Board 2'),
  ('Board 3');