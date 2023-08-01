package models

// type BiEmple struct {
// 	CodEmpl string
// }

type NmContr struct {
	NomEmpl string `gorm:"column:f_name"`
	ApeEmpl string `gorm:"column:l_name"`
}
