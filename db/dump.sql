CREATE DATABASE db;

\c db;

CREATE TABLE users (
	user_id serial PRIMARY KEY,
	name VARCHAR ( 50 ) UNIQUE NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL
);
