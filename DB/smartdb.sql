	CREATE TABLE Users (
	id SERIAL PRIMARY KEY,
	full_name VARCHAR(50),
	email VARCHAR(100),
	pass VARCHAR(16)
);

CREATE TABLE roles (
	id SERIAL PRIMARY KEY,
	rol VARCHAR(25)
);

ALTER TABLE users ADD COLUMN fk_rol INT;

ALTER TABLE users
ADD CONSTRAINT fk_rol
FOREIGN KEY(fk_rol) 
REFERENCES roles(id);

CREATE TABLE pqrs_sac(
	id SERIAL PRIMARY KEY,
	full_name VARCHAR(50) not null,
	fk_document_type integer not null,
	document_id integer not null,
	regional VARCHAR(15),
	program_name VARCHAR(10) not null,
	campus VARCHAR(25),
	languaje VARCHAR(9),
	fk_pqrs_type_id integer,
	descreption_msg VARCHAR(255)
);

CREATE TABLE pqrs_type(
	id SERIAL PRIMARY KEY,
	pqrs_type VARCHAR(25)
);

ALTER TABLE pqrs_sac 
ADD CONSTRAINT fk_pqrs_type_id 
FOREIGN KEY(fk_pqrs_type_id)
REFERENCES pqrs_type(id)

CREATE TABLE document_type(
	id SERIAL PRIMARY KEY,
	document_type VARCHAR(25)
);

ALTER TABLE pqrs_sac 
ADD CONSTRAINT fk_document_type 
FOREIGN KEY(fk_document_type)
REFERENCES document_type(id)

/* 
/////////////////////////////////  */

insert into document_type (id, document_type)
values (1, 'cédula de extranjería')

