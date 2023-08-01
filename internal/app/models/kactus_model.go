package models

// type BiEmple struct {
// 	CodEmpl string
// }

type NmContr struct {
	FName    string
	LName    string `gorm:"column:l_name"`
	Position any
	Date     any
	BMail    string
	EMail    string
	FnLeader string
	LnLeader string
}
