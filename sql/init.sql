CREATE TABLE films (
	id VARCHAR(255) PRIMARY KEY,
	name VARCHAR(255),
	year VARCHAR(255),
	rating DOUBLE
	) ENGINE=InnoDB;

CREATE INDEX idx_films ON films (id, name, year, rating);