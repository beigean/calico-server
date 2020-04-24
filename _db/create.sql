-- DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
    created_at datetime  DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    mail varchar(90) NOT NULL UNIQUE,
    password binary(60) NOT NULL,
    name varchar(20),
    age integer
);
INSERT INTO users(name, mail, password, age)
VALUES
  ('Alice', 'alice@sample.com', alice19910101, 20),
  ('Bob', 'bob@sample.com', bob19910201,, 30),
  ('Carol', 'carol@sample.com', carol19910301,, 25);

-- DROP TABLE IF EXISTS todos;
CREATE TABLE todos (
    id integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
    created_at datetime  DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    user_id integer,
    todo varchar(255),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
INSERT INTO todos(user_id, todo)
VALUES
  (1, 'buy tamago'),
  (1, 'wash dish'),
  (3, 'buy hand soap');
