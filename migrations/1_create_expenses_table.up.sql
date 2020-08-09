CREATE TABLE expenses (
  id serial PRIMARY KEY,
	paidAt VARCHAR,
	title VARCHAR ( 100 ) NOT NULL,
	amount INT NOT NULL
);