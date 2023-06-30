package controllers

import (
	"net/http"

	models "github.com/DevEdwinF/smartback.git/internal/app/models/user"
	"github.com/DevEdwinF/smartback.git/internal/config"
	entity "github.com/DevEdwinF/smartback.git/internal/infrastructure/entity/colaborator"
	"github.com/labstack/echo/v4"
)

func GetAllCollaborators(c echo.Context) error {
	collaboratorWithSchedule := []entity.CollaboratorsDataEntity{}

	config.DB.Table("collaborators").Select("*").Joins("left join schedule s on s.id = collaborators.fk_schedule_id").Scan(&collaboratorWithSchedule)
	return c.JSON(http.StatusOK, collaboratorWithSchedule)
}

func GetCollaborator(c echo.Context) error {
	id := c.Param("doc")

	collaboratorWithSchedule := models.CollaboratorsData{}

	config.DB.Table("collaborators").Select("*").Joins("left join schedule s on s.id = collaborators.fk_schedule_id").Where("employes.id = ?", id).First(&collaboratorWithSchedule)

	if collaboratorWithSchedule.Document == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "No se encuentra el colaborador")
	}

	return c.JSON(http.StatusOK, collaboratorWithSchedule)
}

func SaveCollaborator(c echo.Context) error {
	collaborator := entity.CollaboratorsEntity{}

	err := c.Bind(&collaborator)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	collaboratorFromDb := models.Collaborators{}

	config.DB.Table("collaborators").Where("collaborators.document = ?", collaborator.Document).Scan(&collaborator)

	config.DB.Save(&collaboratorFromDb)

	return c.JSON(http.StatusCreated, collaborator)
}

func DeleteCollaborator(c echo.Context) error {
	id := c.Param("doc")

	employee := models.Collaborators{}

	config.DB.Find(&employee, id)

	if employee.Document > 0 {
		config.DB.Delete(employee)
		return c.JSON(http.StatusOK, employee)
	} else {
		return echo.NewHTTPError(http.StatusNotFound, "El colaborador no se encuentra en la base de datos")
	}
}
