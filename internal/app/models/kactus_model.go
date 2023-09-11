package models

// type BiEmple struct {
// 	CodEmpl string
// }

type NmContr struct {
	Document       string `gorm:"column:document"`
	FName          string
	LName          string
	Position       string `gorm:"column:position"`
	Date           any    `gorm:"column:date"`
	State          string
	BMail          string
	EMail          string
	FnLeader       string
	LnLeader       string
	LeaderDocument string
	Subprocess     string `gorm:"column:subproceso"` // Nuevo campo
	Headquarters   string `gorm:"column:sede"`       // Nuevo campo
}
