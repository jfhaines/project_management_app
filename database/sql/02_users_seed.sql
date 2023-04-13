USE project_management_app;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id VARCHAR(40),
  username VARCHAR(60) NOT NULL,
  email VARCHAR(120) NOT NULL,
  password VARCHAR(70) NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO users (username, email, password)
VALUES
  ('user1', 'user1@gmail.com', 'asdf'),
  ('user2', 'user2@gmail.com', 'asdf'),
  ('user3', 'user3@gmail.com', 'asdf');