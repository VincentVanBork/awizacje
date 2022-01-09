CREATE TABLE IF NOT EXISTS users (
    user_id serial PRIMARY KEY,
    email VARCHAR(255) NOT NULL unique,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    password VARCHAR (255) NOT NULL,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS groups
(
    group_id   serial PRIMARY KEY,
    group_name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS user_roles (
   user_id INT NOT NULL,
   role_id INT NOT NULL,
   grant_date TIMESTAMP,
   PRIMARY KEY (user_id, role_id),
   FOREIGN KEY (role_id)
       REFERENCES groups (group_id),
   FOREIGN KEY (user_id)
       REFERENCES users (user_id)
);