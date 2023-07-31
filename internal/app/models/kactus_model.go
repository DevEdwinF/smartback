package models

// type BiEmple struct {
// 	CodEmpl string
// }

type NmContr struct {
	CodEmpl string `gorm:"column:cod_empl"`
	NomEmpl string `gorm:"column:nom_empl"`
	ApeEmpl string `gorm:"column:ape_empl"`
}
