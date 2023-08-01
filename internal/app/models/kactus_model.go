package models

// type BiEmple struct {
// 	CodEmpl string
// }

type NmContr struct {
	Document int `gorm:"column:document"`
	FName    string
	LName    string
	Position any `gorm:"column:position"`
	Date     any `gorm:"column:date"`
	BMail    string
	EMail    string
	FnLeader string
	LnLeader string
}
