CREATE DATABASE db;

\c db;

CREATE TABLE users (
	user_id serial PRIMARY KEY,
	name VARCHAR ( 50 ) UNIQUE NOT NULL,
	email VARCHAR ( 100 ) UNIQUE NOT NULL,
	password VARCHAR ( 255 ) NOT NULL
);
