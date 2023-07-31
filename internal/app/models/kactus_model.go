package models

import "time"

// type BiEmple struct {
// 	CodEmpl string
// }

type Colaborador struct {
	NúmeroIdentificación string    `gorm:"column:cod_inte"`
	Nombres              string    `gorm:"column:nom_empl"`
	Apellidos            string    `gorm:"column:ape_empl"`
	Cargo                string    `gorm:"column:nom_carg"`
	FechaIngreso         time.Time `gorm:"column:fec_ingr"`
	CorreoCorporativo    string    `gorm:"column:box_mail"`
	CorreoPersonal       string    `gorm:"column:eee_mail"`
	NombresJefeInme      string    `gorm:"column:nom_empl"`
	ApellidosJefeInme    string    `gorm:"column:ape_empl"`
}
