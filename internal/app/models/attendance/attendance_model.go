package models

import (
	"time"

	"github.com/DevEdwinF/smartback.git/internal/config"
)

type AttendanceModel struct {
	ID        int64 `json:"document"`
	Arrival   time.Time
	Departure time.Time
}

func GetAllAttendance() ([]AttendanceModel, error) {
	var attendances []AttendanceModel

	err := config.DB.Table("attendance a").Select("c.name, a.* ").Joins("INNER JOIN colaborators e on c.id = a.fk_document_id").Find(&attendances).Error
	if err != nil {
		return nil, err
	}

	return attendances, nil
}

/*
create table Colaborators(
	id serial primary key,
	name varchar(60) not null,
	email varchar(60) not null,
	create_at timestamp
);

create table Attendance(
	id serial primary key,
	arrival timestamp,
	departure timestamp
)

create table Schedule(
	id serial primary key,
	arrival varchar(25Time
	depardure varchar(25)
)

ALTER TABLE Colaborators
add column fk_schedule_id integer,
ADD CONSTRAINT fk_schedule_id
FOREIGN KEY (fk_schedule_id)
REFERENCES Schedule (id);

ALTER TABLE attendance
add column fk_attendance_id integer,
ADD CONSTRAINT fk_attendance_id
FOREIGN KEY (fk_attendance_id)
REFERENCES Attendance (id);
*/
