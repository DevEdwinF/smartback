
create table collaborators(
    document int primary key,			
	name varchar(50) not null,
	email varchar(50) not null,
	position varchar(45) not null,
	leader varchar(50) not null,
	create_at timestamp
);

create table attendances(
	id serial primary key,
	arrival timestamp,
	departure timestamp,
	location varchar(10),
	photo bytea,
    created_at timestamp
)

ALTER TABLE attendances
ADD COLUMN fk_document_id int,
ADD CONSTRAINT fk_document_id
FOREIGN KEY (fk_document_id)
REFERENCES collaborators (document);

INSERT INTO "collaborators" ("document", "name", "email", "position", "leader")
VALUES (1032500648, 'Edwin Fernando Pirajan Arevalo', 'epiraja@smart.edu.co', 'Desarrollador de software', 'Jorge Celemin');

create table schedule(
	id serial primary key,
	day varchar(11),
	arrival_time TIMESTAMP,
	departure_time TIMESTAMP
)

ALTER TABLE schedule
add column fk_collaborators_document integer,
ADD CONSTRAINT fk_collaborators_document
FOREIGN KEY (fk_collaborators_document)
REFERENCES collaborators(document);

INSERT INTO "schedule" ("day", "arrival_time", "departure_time")
VALUES ('Monday', '07:00:00', '17:00:00');




SELECT * FROM "attendances" WHERE fk_document_id = 123 AND DATE(created_at) = CURRENT_DATE ORDER BY "attendances"."id" LIMIT 1