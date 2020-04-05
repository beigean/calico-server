-- DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name varchar(20),
    age integer
);
INSERT INTO users(name, age)
VALUES
  ('Alice', 20),
  ('Bob', 30),
  ('Carol', 40);

-- DROP TABLE IF EXISTS todos;
CREATE TABLE todos (
    id integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user_id integer,
    todo varchar(255),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
INSERT INTO todos(user_id, todo)
VALUES
  (1, 'buy tamago'),
  (1, 'wash dish'),
  (3, 'buy hand soap');
