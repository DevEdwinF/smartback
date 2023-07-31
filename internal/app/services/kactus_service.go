package services

import (
	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
)

func GetAllColab() ([]models.Colaborador, error) {
	collaborators := []models.Colaborador{}
	err := config.DB.Raw(`
        SELECT
            bi_emple.cod_inte AS NúmeroIdentificación,
            bi_emple.nom_empl AS Nombres,
            bi_emple.ape_empl AS Apellidos,
            nom_carg AS Cargo,
            fec_ingr AS FechaIngreso,
            bi_emple.box_mail AS CorreoCorporativo,
            bi_emple.eee_mail AS CorreoPersonal,
            NomJefe.nom_empl AS NombresJefeInme,
            NomJefe.ape_empl AS ApellidosJefeInme
        FROM
            NM_CONTR
        INNER JOIN
            bi_emple ON NM_CONTR.cod_empl = bi_emple.cod_empl
        INNER JOIN
            BI_CARGO ON NM_CONTR.cod_carg = BI_CARGO.cod_carg
        INNER JOIN
            bi_emple AS NomJefe ON NM_CONTR.cod_frep = NomJefe.cod_empl
    `).Scan(&collaborators).Error
	return collaborators, err
}
