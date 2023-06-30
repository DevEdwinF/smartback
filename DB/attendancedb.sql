drop table collaborators , attendances, schedule

create table collaborators(
    document int primary key,			
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

ALTER TABLE attendances
ADD COLUMN fk_document_id int,
ADD CONSTRAINT fk_document_id
FOREIGN KEY (fk_document_id)
REFERENCES collaborators (document);

INSERT INTO "collaborators" ("document", "name", "email")
VALUES (1032500648, 'Edwin Fernando Pirajan Arevalo', 'epiraja@smart.edu.co');

create table schedule(
	id serial primary key,
	arrival varchar(25),
	depardure varchar(25)
)

ALTER TABLE collaborators
add column fk_schedule_id integer
ADD CONSTRAINT fk_schedule_id
FOREIGN KEY (fk_schedule_id)
REFERENCES Schedule (id);




SELECT * FROM "attendances" WHERE fk_document_id = 123 AND DATE(created_at) = CURRENT_DATE ORDER BY "attendances"."id" LIMIT 1