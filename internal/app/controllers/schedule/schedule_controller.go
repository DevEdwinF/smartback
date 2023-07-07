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

// func saveSchedule(schedule entity.ScheduleEntity) (models.ScheduleModel, error) {
// 	scheduleV := models.ScheduleModel{}
// 	scheduleV.Id = schedule.Id
// 	scheduleV.Day = schedule.Day
// 	scheduleV.ArrivalTime = schedule.ArrivalTime
// 	scheduleV.DepartureTime = schedule.DepartureTime
// 	scheduleV.FkDocument = schedule.FkDocument

// 	err := config.DB.Save(&scheduleV).Error
// 	if err != nil {
// 		return models.ScheduleModel{}, err
// 	}
// 	return scheduleV, nil
// }

// func GetAllSchedule(c echo.Context) error {

// 	schedule := []entity.ScheduleEntity{}
// 	config.DB.Find(&schedule)
// 	return c.JSON(http.StatusOK, schedule)
// }

// func SaveSchedule(c echo.Context) error {
// 	requestBody := entity.ScheduleEntity{}
// 	if err := c.Bind(&requestBody); err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}

// 	schedule, err := saveSchedule(requestBody)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}
// 	return c.JSON(http.StatusCreated, schedule)
// }

func GetAllCollaboratorsSchedule(c echo.Context) error {
	collaboratorWithSchedule := []entityCollaborator.CollaboratorsDataEntity{}

	config.DB.Table("collaborators").Select("*").
		Joins("left join schedule on collaborators.document = schedule.fk_collaborators_document").
		Scan(&collaboratorWithSchedule)

	return c.JSON(http.StatusOK, collaboratorWithSchedule)
}

func SaveSchedule(c echo.Context) error {
	var schedule entity.ScheduleEntity
	if err := c.Bind(&schedule); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Formato de datos inválido")
	}

	result := config.DB.Table("schedule").Create(&schedule)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error al guardar el horario")
	}

	return c.JSON(http.StatusOK, schedule)
}

func EditSchedule(c echo.Context) error {
	idParam := c.Param("id")
	var schedule entity.ScheduleEntity

	if err := config.DB.Table("schedule").First(&schedule, "id = ?", idParam).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "No se encuentra el horario")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Error en el servidor")
	}

	updatedSchedule := entity.ScheduleEntity{}
	if err := c.Bind(&updatedSchedule); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Formato de datos inválido")
	}

	result := config.DB.Model(&schedule).Updates(updatedSchedule)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error al actualizar el horario")
	}

	return c.JSON(http.StatusOK, schedule)
}

func AssignScheduleToCollaborator(c echo.Context) error {
	var schedule entity.ScheduleEntity

	if err := c.Bind(&schedule); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Formato de datos inválido")
	}

	// Verificamos que el colaborador exista
	var collaborator entityCollaborator.CollaboratorsDataEntity
	if err := config.DB.Table("collaborators").Take(&collaborator, "document = ?", schedule.FkCollaboratorsDocument).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "No se encuentra el colaborador")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Error en el servidor")
	}

	// Agregamos el horario
	result := config.DB.Table("schedule").Create(&schedule)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error al asignar el horario")
	}

	return c.JSON(http.StatusOK, schedule)
}

func DeleteSchedule(c echo.Context) error {
	id := entity.ScheduleEntity{}

	if err := c.Bind(&id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var schedule models.ScheduleModel
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
