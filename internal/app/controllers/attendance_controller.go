package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
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

	var schedule models.Schedules
	err = config.DB.Model(&schedule).Where("fk_collaborators_document = ? AND day = ?", attendance.FkDocumentId, time.Now().Format("Monday")).First(&schedule).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	timeNow := time.Now()

	var arrivalScheduled time.Time
	if schedule.ArrivalTime != "" {
		temp, err := time.Parse("15:04:05", schedule.ArrivalTime)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		arrivalScheduled = time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), temp.Hour(), temp.Minute(), temp.Second(), temp.Nanosecond(), timeNow.Location())
	}

	late := false

	fmt.Println("Horario asignado", arrivalScheduled)
	fmt.Println("tiempo actual", timeNow)

	if !arrivalScheduled.IsZero() && timeNow.After(arrivalScheduled.Add(5*time.Minute)) {
		fmt.Println("entra?")
		late = true
	}

	var validateAttendance models.Attendance
	err = config.DB.Model(&validateAttendance).
		Where("fk_document_id = ? AND DATE(created_at) = ?", attendance.FkDocumentId, timeNow.Format("2006-01-02")).
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
			FkDocumentId: attendance.FkDocumentId,
			Photo:        attendance.Photo,
			Location:     attendance.Location,
			Arrival:      sql.NullString{String: timeNow.Format("15:04:05"), Valid: true},
			Late:         late,
			CreatedAt:    timeNow,
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

	config.DB.Table("attendances a").Select("c.f_name, c.l_name,c.email, a.* ").Joins("INNER JOIN collaborators c on c.document = a.fk_document_id").Find(&attendance)

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

func SaveTranslated(c echo.Context) error {
	var translatedEntity entity.Translatedcollaborators
	if err := c.Bind(&translatedEntity); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err := ValidateCollaborator(translatedEntity.FkDocumentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	translatedEntity.CreatedAt = time.Now()

	if err := config.DB.Create(&translatedEntity).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Registro de traducción creado exitosamente",
	})
}

func ValidateColaborator(c echo.Context) error {
	id := c.Param("doc")

	var employe models.Collaborators
	if err := config.DB.Table("collaborators").Where("document = ?", id).Scan(&employe).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if employe.Document == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Empleado no se encuentra registrado")
	}
	return c.JSON(http.StatusOK, employe)
}

func ValidateCollaborator(document int) (*models.Collaborators, error) {
	var collaborator models.Collaborators
	if err := config.DB.Table("collaborators").Where("document = ?", document).Scan(&collaborator).Error; err != nil {
		return nil, err
	}
	if collaborator.Document == 0 {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Empleado no se encuentra registrado")
	}
	return &collaborator, nil
}
