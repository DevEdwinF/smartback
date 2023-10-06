create table
    collaborators(
        id serial primary key,
        document VARCHAR(25) UNIQUE,
        f_name varchar(50) not null,
        l_name varchar(50) not null,
        email varchar(100) not null,
        bmail varchar(100) not null,
        position varchar(45) not null,
        state VARCHAR(10) not null,
        leader varchar(50) not null,
        leader_document VARCHAR(20),  
        process VARCHAR(50),
        subprocess VARCHAR(50),
        headquarters VARCHAR(30),
        created_at timestamp
    );

create table
    attendances(
        id serial primary key,
        arrival time,
        departure time,
        location varchar(10),
        late BOOLEAN,
        early_arrival BOOLEAN,
        photo_arrival bytea,
        photo_departure bytea,
        created_at timestamp
    )

ALTER TABLE attendances
ADD
    COLUMN fk_collaborator_id int,
ADD
    CONSTRAINT fk_collaborator_id FOREIGN KEY (fk_collaborator_id) REFERENCES collaborators (id);


create table
    schedules(
        id serial primary key,
        day varchar(11),
        arrival_time VARCHAR,
        departure_time VARCHAR
    )

ALTER TABLE schedules
add
    column fk_collaborator_id integer,
ADD
    CONSTRAINT fk_collaborator_id FOREIGN KEY (fk_collaborator_id) REFERENCES collaborators(id);

create table
    TranslatedCollaborators (
        id serial primary key,
        created_at timestamp
    );

ALTER TABLE
    TranslatedCollaborators
add
    column fk_collaborator_id integer,
ADD
    CONSTRAINT fk_collaborator_id FOREIGN KEY (fk_collaborator_id) REFERENCES collaborators(id);

create table
    Users (
        id serial primary key,
        document VARCHAR(25),
        f_name varchar(50) not null,
        l_name varchar(50) not null,
        email varchar(100) not null,
        password varchar(12) not null,
        created_at timestamp
    )

create table roles ( id serial primary key, name varchar(25) ) 

ALTER TABLE users
ADD
    COLUMN fk_role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE;

INSERT INTO
    "users" (
        "f_name",
        "l_name",
        "email",
        "fk_role_id",
        "password"
    )
VALUES (
        'Edwin Fernando',
        'Pirajan Arevalo',
        'epirajan@smart.edu.co',
        1,
        'Smart2023++'
    );

CREATE TABLE headquarters (
        id serial primary key,
        name varchar(50) not null
    )

INSERT INTO "headquarters" ("name")
VALUES ('ADMINISTRATIVO'), ('ARKADIA'), ('BELLO'), ('CALASANZ'), ('CALIMA'), ('CEDRITOS'), ('CENTRO INTERNACIONAL'), ('CENTRO MAYOR'), ('CENTRO MEDELLIN'), ('CHAPINERO'), ('CHIA'), ('ENVIGADO'), ('FLORIDABLANCA'), ('FONTANAR'), ('HAYUELOS'), ('ITAGÜÍ'), ('MADELENA'), ('MODELIA'), ('MULTIPLAZA'), ('NUESTRO BOGOTÁ'), ('OLAYA'), ('PALATINO'), ('PIEDECUESTA'), ('PLAZA CENTRAL'), ('PLAZA DE LAS AMERICAS'), ('POBLADO'), ('RESTREPO'), ('SAN MARTÍN'), ('SANTAFÉ'), ('ONLINE'), ('SOACHA'), ('SUBA'), ('UNICENTRO DE OCCIDENTE'), ('VIRTUAL');