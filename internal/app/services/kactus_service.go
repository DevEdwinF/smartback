package services

import (
	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
)

func GetAllColab() ([]models.NmContr, error) {
	collaborators := []models.NmContr{}
	err := config.KDB.Table("NM_CONTR").
		Select(`
            bi_emple.cod_inte AS document,
            bi_emple.nom_empl AS f_name,
            bi_emple.ape_empl AS l_name,
            BI_CARGO.nom_carg AS position,
            NM_CONTR.fec_ingr AS date,
            bi_emple.box_mail AS b_mail,
            bi_emple.eee_mail AS e_mail,
            NomJefe.cod_inte AS "Jefe documento",
            NomJefe.nom_empl AS fn_leader,
            NomJefe.ape_empl AS ln_leader,
            gc.nom_ccos AS subprocess,
            sa.NOM_AREA AS headquarters,
            NM_CONTR.ind_acti AS state`).
		Joins(`
            INNER JOIN bi_emple ON NM_CONTR.cod_empl = bi_emple.cod_empl
            INNER JOIN BI_CARGO ON NM_CONTR.cod_carg = BI_CARGO.cod_carg
            INNER JOIN bi_emple AS NomJefe ON NM_CONTR.cod_frep = NomJefe.cod_empl
            LEFT JOIN gn_ccost gc ON NM_CONTR.cod_ccos = gc.cod_ccos
            LEFT JOIN SO_AREAS sa ON NM_CONTR.cod_area = sa.COD_AREA`).
		Where("NM_CONTR.ind_acti = ?", "A").
		Scan(&collaborators).Error
	if err != nil {
		return nil, err
	}
	return collaborators, nil
}

func GetColabById(id string) (*models.NmContr, error) {
	collaborator := models.NmContr{}
	err := config.KDB.Table("NM_CONTR").
		Select(`
            bi_emple.cod_inte AS document,
            bi_emple.nom_empl AS f_name,
            bi_emple.ape_empl AS l_name,
            BI_CARGO.nom_carg AS position,
            NM_CONTR.fec_ingr AS date,
            bi_emple.box_mail AS b_mail,
            bi_emple.eee_mail AS e_mail,
            NomJefe.nom_empl AS fn_leader,
            NomJefe.ape_empl AS ln_leader,
            NM_CONTR.ind_acti AS state`).
		Joins(`
            INNER JOIN bi_emple ON NM_CONTR.cod_empl = bi_emple.cod_empl
            INNER JOIN BI_CARGO ON NM_CONTR.cod_carg = BI_CARGO.cod_carg
            INNER JOIN bi_emple AS NomJefe ON NM_CONTR.cod_frep = NomJefe.cod_empl`).
		Where("bi_emple.cod_inte = ?", id).
		Scan(&collaborator).Error

	if err != nil {
		return nil, err
	}

	return &collaborator, nil
}
