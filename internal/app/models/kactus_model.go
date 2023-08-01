package models

// type BiEmple struct {
// 	CodEmpl string
// }

type NmContr struct {
	FName    string
	ApeEmpl  string `gorm:"column:l_name"`
	Position any    `gorm:"column:position"`
	Date     any    ``
}
