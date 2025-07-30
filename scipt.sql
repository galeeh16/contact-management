-- create table users
CREATE TABLE IF NOT EXISTS users( 
	id SERIAL PRIMARY KEY,
	username VARCHAR(255) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL,
	name VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP
);

-- create table contacts
CREATE TABLE IF NOT EXISTS contacts(
	id SERIAL PRIMARY KEY,
	first_name VARCHAR(100) NOT NULL,
	last_name VARCHAR(255),
	phone VARCHAR(20) NOT NULL,
	user_id INTEGER NOT NULL,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id)
);

-- password = Secret
INSERT INTO users(username, password, name) VALUES ('username1', '$2a$12$8MVykBX6mGn0eVNM0XlTQORGfHDGqGjOf4l6Wg31y9jvSEpRN7wnC', 'Nama User Satu');

-- DROP TABLE contacts;
-- DROP TABLE users;