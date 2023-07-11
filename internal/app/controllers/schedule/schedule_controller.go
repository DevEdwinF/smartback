package controllers

import (
	"errors"
	"net/http"

	models "github.com/DevEdwinF/smartback.git/internal/app/models/schedule"
	"github.com/DevEdwinF/smartback.git/internal/config"
	entityCollaborator "github.com/DevEdwinF/smartback.git/internal/infrastructure/entity/colaborator"
	entity "github.com/DevEdwinF/smartback.git/internal/infrastructure/entity/schedule"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func GetAllCollaboratorsSchedule(c echo.Context) error {
	collaboratorWithSchedule := []entityCollaborator.CollaboratorsDataEntity{}

	config.DB.Table("collaborators").Select("*").
		Joins("left join schedule on collaborators.document = schedule.fk_collaborators_document").
		Scan(&collaboratorWithSchedule)

	return c.JSON(http.StatusOK, collaboratorWithSchedule)
}

func AssignSchedulesToCollaborator(c echo.Context) error {
	var schedules []entity.Schedule

	if err := c.Bind(&schedules); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Formato de datos inválido")
	}

	// Verificamos que el colaborador exista
	var collaborator entityCollaborator.CollaboratorsDataEntity
	if err := config.DB.Table("collaborators").Take(&collaborator, "document = ?", schedules[0].FkCollaboratorsDocument).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "No se encuentra el colaborador")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Error en el servidor")
	}

	for _, schedule := range schedules {
		// Buscamos el horario existente o creamos uno nuevo
		var existingSchedule entity.Schedule
		result := config.DB.Table("schedule").Where("id = ?", schedule.Id).Assign(schedule).FirstOrCreate(&existingSchedule)

		if result.Error != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error al asignar el horario")
		}
	}

	return c.JSON(http.StatusOK, schedules)
}

func DeleteSchedule(c echo.Context) error {
	id := entity.Schedule{}

	if err := c.Bind(&id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var schedule models.Schedule
	if err := config.DB.First(&schedule, id.Id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Schedule not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Delete(&schedule).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Schedule deleted")
}
