package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AttendanceController struct {
	Service *services.AttendanceService
}

func NewAttendanceController(service *services.AttendanceService) *AttendanceController {
	return &AttendanceController{
		Service: service,
	}
}

func (ac *AttendanceController) SaveRegisterAttendance(c echo.Context) error {
	var attendance entity.AttendanceEntity
	err := c.Bind(&attendance)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = ac.Service.RegisterAttendance(attendance)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Registro de asistencia guardado exitosamente",
	})
}

func (controller *AttendanceController) GetAllAttendance(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
	}

	pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))
	if err != nil {
		pageSize = 100
	}

	attendance, err := controller.Service.GetAttendancePage(page, pageSize)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, attendance)
}

func (controller *AttendanceController) GetAllAttendanceForLate(c echo.Context) error {
	attendance, err := controller.Service.GetAllAttendanceForToLate()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, attendance)
}

func (controller *AttendanceController) GetAttendanceForLeader(c echo.Context) error {

	userToken := c.Get("userToken")

	if userToken == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token de usuario no encontrado")
	}

	token, ok := userToken.(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error al procesar el token")
	}

	claims := token.Claims.(jwt.MapClaims)

	leaderDocument, ok := claims["document"].(string)

	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"error": "Este usuario no tiene ningún documento de líder asignado",
		})
	}

	attendanceService := &services.AttendanceService{}
	attendance, err := attendanceService.GetAttendanceForLeader(leaderDocument)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Error obteniendo la asistencia",
		})
	}

	return c.JSON(http.StatusOK, attendance)
}

func (controller *AttendanceController) GetAttendanceForLeaderToLate(c echo.Context) error {
	userToken := c.Get("userToken")

	if userToken == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token de usuario no encontrado")
	}

	token, ok := userToken.(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error al procesar el token")
	}

	claims := token.Claims.(jwt.MapClaims)

	leaderDocument, ok := claims["document"].(string)

	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"error": "Este usuario no tiene ningún documento de líder asignado",
		})
	}

	attendanceService := &services.AttendanceService{}
	attendance, err := attendanceService.GetAttendanceForLeaderToLate(leaderDocument)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Error obteniendo la asistencia",
		})
	}

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

	collaborator, err := services.GetCollaborator(document)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, collaborator)
}

func (ac *AttendanceController) SaveTranslated(c echo.Context) error {
	var translated entity.Translatedcollaborators
	err := c.Bind(&translated)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = ac.Service.SaveTranslatedService(translated)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Registro de asistencia guardado exitosamente",
	})
}

func GetAllTranslatedController(c echo.Context) error {
	translatedcollaborators, err := services.GetAllTranslatedService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, translatedcollaborators)
}
