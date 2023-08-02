package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SaveRegisterAttendance(c echo.Context) error {
	var attendance entity.AttendanceEntity
	err := c.Bind(&attendance)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var collaborator models.Collaborators
	err = config.DB.Model(&collaborator).Where("document = ?", attendance.Document).First(&collaborator).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return echo.NewHTTPError(http.StatusNotFound, "Colaborador no encontrado")
	}

	timeNow := time.Now()

	var schedule models.Schedules
	err = config.DB.Model(&schedule).Where("fk_collaborator_id = ? AND day = ?", collaborator.Id, timeNow.Format("Monday")).First(&schedule).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return echo.NewHTTPError(http.StatusNotFound, "Horario no encontrado para el colaborador en este día")
	}

	var arrivalScheduled time.Time
	if schedule.ArrivalTime != "" {
		temp, err := time.Parse("15:04:05", schedule.ArrivalTime)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		arrivalScheduled = time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), temp.Hour(), temp.Minute(), temp.Second(), temp.Nanosecond(), timeNow.Location())
	}

	late := false

	if !arrivalScheduled.IsZero() && timeNow.After(arrivalScheduled.Add(5*time.Minute)) {
		late = true
	}

	var validateAttendance models.Attendance
	err = config.DB.Model(&validateAttendance).
		Where("fk_collaborator_id = ? AND date(created_at) = ?", collaborator.Id, timeNow.Format("2006-01-02")).
		First(&validateAttendance).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	switch attendance.State {
	case "arrival":
		if validateAttendance.ID != 0 || validateAttendance.Arrival.Valid {
			return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado la entrada")
		}

		modelsAttendance := models.Attendance{
			FkCollaboratorID: collaborator.Id,
			Photo:            attendance.Photo,
			Location:         attendance.Location,
			Arrival:          sql.NullString{String: timeNow.Format("15:04:05"), Valid: true},
			Late:             late,
			CreatedAt:        timeNow,
		}
		err = config.DB.Create(&modelsAttendance).Error
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Registro de entrada creado exitosamente",
		})

	case "departure":
		if validateAttendance.ID == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "Debe registrar la entrada primero")
		}
		if validateAttendance.Departure.Valid {
			return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado la salida")
		}

		validateAttendance.Departure = sql.NullString{String: timeNow.Format("15:04:05"), Valid: true}

		err = config.DB.Updates(&validateAttendance).Error
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Registro de salida actualizado exitosamente",
		})
	}

	return echo.NewHTTPError(http.StatusBadRequest, "Estado inválido")
}

func GetAllAttendance(c echo.Context) error {
	attendance := []entity.UserAttendanceData{}

	config.DB.Table("attendances a").Select("c.f_name, c.l_name, c.email, c.document, a.* ").Joins("INNER JOIN collaborators c on c.id = a.fk_collaborator_id").Find(&attendance)

	return c.JSON(http.StatusOK, attendance)
}

func ValidateSchedule(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var validateSchedule entity.ValidateSchedule
	err = json.Unmarshal(body, &validateSchedule)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var arrival time.Time

	config.DB.Raw("select arrival from attendance a where fk_document_id = ? and date_format(arrival, '%d-%m-%Y') = date_format(?, '%d-%m-%Y')",
		validateSchedule.Id, validateSchedule.Date).Scan(&arrival)

	return c.JSON(http.StatusOK, arrival)
}

func ValidateCollaboratorController(c echo.Context) error {
	document := c.Param("doc")

	collaborator, err := services.ValidateCollaboratorService(document)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, collaborator)
}

func SaveTranslated(c echo.Context) error {
	var translatedEntity entity.Translatedcollaborators
	err := c.Bind(&translatedEntity)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Buscar el colaborador por el documento
	var collaborator models.Collaborators
	err = config.DB.Model(&collaborator).Where("document = ?", translatedEntity.Document).First(&collaborator).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return echo.NewHTTPError(http.StatusNotFound, "Colaborador no encontrado")
	}

	newTranslatedCollaborator := models.Translatedcollaborators{
		FkCollaboratorId: collaborator.Id,
		CreatedAt:        time.Now(),
	}

	err = config.DB.Create(&newTranslatedCollaborator).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Translado registrado con éxito",
	})
}
