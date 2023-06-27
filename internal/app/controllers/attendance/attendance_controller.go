package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	models "github.com/DevEdwinF/smartback.git/internal/app/models/attendance"
	modelsColaborator "github.com/DevEdwinF/smartback.git/internal/app/models/user"
	"github.com/DevEdwinF/smartback.git/internal/config"
	entity "github.com/DevEdwinF/smartback.git/internal/infrastructure/entity/attendance"
	"github.com/labstack/echo/v4"
)

func getAllAttendanceController() {

}

func SaveRegisterAttendance(c echo.Context) error {
	var attendance entity.AttendanceEntity
	err := c.Bind(&attendance)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var validateAttendance models.Attendance
	if err := config.DB.Model(&validateAttendance).Where("fk_document_id = ? AND DATE(created_at) = CURRENT_DATE", attendance.FkDocumentId).Scan(&validateAttendance).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Ops, este documento no se encuentra registrado")
	}
	//location, _ := time.LoadLocation("America/Bogota")
	timeNow := time.Now() /*.In(location)*/

	fmt.Println(timeNow)

	if attendance.State == "arrival" {
		if validateAttendance.ID == 0 && validateAttendance.Arrival == nil {
			modelsAttendance := models.Attendance{
				FkDocumentId: attendance.FkDocumentId,
				// Photo:        attendance.Photo,
				Arrival:   &timeNow,
				CreatedAt: timeNow,
			}

			err = config.DB.Create(&modelsAttendance).Error
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return c.JSON(http.StatusOK, map[string]string{
				"message": "Registro creado exitosamente",
			})
		}
	}

	block := validateAttendance.Arrival == nil
	// block := validateAttendance.Departure == nil

	switch attendance.State {
	case "arrival":
		if validateAttendance.Arrival != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado el estado entrada")
		}
		break
	// case "transfer":
	// 	if block {
	// 		return echo.NewHTTPError(http.StatusBadRequest, "Debe registrar la llegada primero")
	// 	}
	// 	break
	case "departure":
		if block {
			return echo.NewHTTPError(http.StatusBadRequest, "Debe registrar la llegada primero")
		}
		if validateAttendance.Departure != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Ya se ha registrado el estado salida")
		}
		validateAttendance.Departure = &timeNow
		break
	}

	err = config.DB.Updates(&validateAttendance).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Registro actualizado exitosamente",
	})
}

func GetAllAttendance(c echo.Context) error {

	attendance := []models.Attendance{}

	config.DB.Table("attendance a").Select("c.name, a.* ").Joins("INNER JOIN colaborators e on c.id = a.fk_document_id").Find(&attendance)

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
	id := c.Param("document")

	var employe modelsColaborator.Colaborators
	if err := config.DB.Table("colaborators").Where("id = ?", id).Scan(&employe).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if employe.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Empleado no se encuentra registrado")
	}
	return c.JSON(http.StatusOK, employe)

}
