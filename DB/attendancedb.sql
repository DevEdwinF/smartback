create table colaborators(
    document int primary key
	name varchar(60) not null,
	email varchar(60) not null,
	create_at timestamp
);

create table attendances(
	id serial primary key,
	arrival timestamp,
	departure timestamp,
    created_at timestamp
)

-- create table schedule(
-- 	id serial primary key,
-- 	arrival varchar(25),
-- 	depardure varchar(25)
-- )

ALTER TABLE attendances
add column fk_document_id int
ADD CONSTRAINT fk_document_id
FOREIGN KEY (fk_document_id)
REFERENCES colaborators (document);

-- ALTER TABLE colaborators
-- add column fk_schedule_id integer
-- ADD CONSTRAINT fk_schedule_id
-- FOREIGN KEY (fk_schedule_id)
-- REFERENCES Schedule (id);
