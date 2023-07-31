package services

import (
	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
)

func GetAllColab() ([]models.BiEmple, error) {
	collaborators := []models.BiEmple{}
	err := config.DB.Raw("SELECT bi_emple.cod_INTE AS NúmeroIdentificación, bi_emple.nom_carg as Cargo, bi_emple.nom_empl as nombres, bi_emple.ape_empl as apellidos, bi_emple.fec_ingr as FechaIngreso, NM_CONTR.box_mail FROM NM_CONTR INNER JOIN bi_emple on NM_CONTR.cod_empl = bi_emple.cod_empl").Scan(&collaborators).Error
	return collaborators, err
}
