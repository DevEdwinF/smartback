package models

// type BiEmple struct {
// 	CodEmpl string
// }

type NmContr struct {
	NomEmpl string `gorm:"column:nom_empl"`
	ApeEmpl string `gorm:"column:ape_empl"`
}
