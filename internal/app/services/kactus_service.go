package services

import (
	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
)

func GetAllColab() ([]models.Collaborators, error) {
	collaborators := []models.Collaborators{}
	err := config.KDB.Table("NM_CONTR").
		Select("bi_emple.cod_inte AS NúmeroIdentificación, bi_emple.nom_empl AS Nombres, bi_emple.ape_empl AS Apellidos, BI_CARGO.nom_carg AS Cargo, NM_CONTR.fec_ingr AS FechaIngreso, bi_emple.box_mail AS CorreoCorporativo, bi_emple.eee_mail AS CorreoPersonal, NomJefe.nom_empl AS NombresJefeInme, NomJefe.ape_empl AS ApellidosJefeInme").
		Joins("INNER JOIN bi_emple ON NM_CONTR.cod_empl = bi_emple.cod_empl").
		Joins("INNER JOIN BI_CARGO ON NM_CONTR.cod_carg = BI_CARGO.cod_carg").
		Joins("INNER JOIN bi_emple AS NomJefe ON NM_CONTR.cod_frep = NomJefe.cod_empl").
		Scan(&collaborators).Error

	if err != nil {
		return nil, err
	}

	return collaborators, nil
}
