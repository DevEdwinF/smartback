package entity

import "time"

// type AttendanceEntity struct {
// 	ID         int64  `json:"document"`
// 	Name       string `json:"name"`
// 	Position   string `json:"position"`
// 	ScheduleId int    `json:"schedule_id"`
// }

type AttendanceEntity struct {
	ID               int64     `json:"id"`
	FkCollaboratorId int       `json:"fk_collaborator_id"`
	Document         string    `json:"document" form:"document"`
	State            string    `json:"state" form:"state"`
	Location         string    `josn:"location" form:"location"`
	Late             *bool     `json:"late"`
	Photo            string    `json:"photo" form:"photo"`
	CreatedAt        time.Time `json:"date"`
}

type UserAttendanceData struct {
	// FkDocumentId int       `json:"document"`
	FkCollaboratorId int       `json:"fk_collaborator_id"`
	Document         string    `json:"document" query:"document"`
	FName            string    `json:"f_name" query:"f_name"`
	LName            string    `json:"l_name" query:"l_name"`
	Email            string    `json:"email" query:"email"`
	Location         string    `json:"location" query:"location"`
	Arrival          string    `json:"arrival" query:"arrival"`
	Departure        string    `json:"departure" query:"departure"`
	Leader           string    `json:"leader" query:"leader"`
	LeaderDocument   string    `json:"leader_document" query:"leader_document"`
	Subprocess       string    `json:"sub_process" query:"sub_process"`
	Late             *bool     `json:"late" query:"late"`
	PhotoArrival     string    `json:"photo_arrival"`
	PhotoDeparture   string    `json:"photo_departure"`
	CreatedAt        time.Time `json:"date"`
}

type AttendanceFilter struct {
	UserAttendanceData
	Paginate
}

type ValidateSchedule struct {
	Id   string `json:"document"`
	Date string `json:"date"`
}

type Translatedcollaborators struct {
	FkCollaboratorId int    `json:"id"`
	Document         string `json:"document"`
	FName            string `json:"f_name"`
	LName            string `json:"l_name"`
	CreatedAt        string `json:"date"`
}
