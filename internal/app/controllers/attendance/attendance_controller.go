package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	models "github.com/DevEdwinF/smartback.git/internal/app/models/attendance"
	modelsColaborator "github.com/DevEdwinF/smartback.git/internal/app/models/user"
	"github.com/DevEdwinF/smartback.git/internal/config"
	entity "github.com/DevEdwinF/smartback.git/internal/infrastructure/entity/attendance"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func getAllAttendanceController() {

}

func SaveRegisterAttendance(c echo.Context) error {
	var attendance entity.AttendanceEntity
	err := c.Bind(&attendance)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	timeNow := time.Now()

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
		if validateAttendance.ID != 0 || validateAttendance.Arrival != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado la entrada")
		}

		modelsAttendance := models.Attendance{
			FkDocumentId: attendance.FkDocumentId,
			Photo:        attendance.Photo,
			Location:     attendance.Location,
			Arrival:      &timeNow,
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
		if validateAttendance.Departure != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado la salida")
		}

		validateAttendance.Departure = &timeNow

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

	config.DB.Table("attendances a").Select("c.name,c.email, a.* ").Joins("INNER JOIN collaborators c on c.document = a.fk_document_id").Find(&attendance)

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

func ValidateColaborator(c echo.Context) error {
	id := c.Param("doc")

	var employe modelsColaborator.Collaborators
	if err := config.DB.Table("collaborators").Where("document = ?", id).Scan(&employe).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if employe.Document == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Empleado no se encuentra registrado")
	}
	return c.JSON(http.StatusOK, employe)
}
