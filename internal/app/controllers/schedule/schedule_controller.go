package controllers

import (
	"net/http"

	models "github.com/DevEdwinF/smartback.git/internal/app/models/schedule"
	"github.com/DevEdwinF/smartback.git/internal/config"
	entity "github.com/DevEdwinF/smartback.git/internal/infrastructure/entity/schedule"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func saveSchedule(schedule entity.ScheduleEntity) (models.ScheduleModel, error) {

	scheduleV := models.ScheduleModel{}

	scheduleV.Id = schedule.Id

	scheduleV.Arrival = schedule.Arrival

	scheduleV.Departure = schedule.Departure

	err := config.DB.Save(&schedule).Error
	if err != nil {
		return models.ScheduleModel{}, err
	}
	return scheduleV, nil
}

func GetAllSchedule(c echo.Context) error {

	schedule := []entity.ScheduleEntity{}
	config.DB.Find(&schedule)
	return c.JSON(http.StatusOK, schedule)
}

func SaveSchedule(c echo.Context) error {
	requestBody := entity.ScheduleEntity{}
	if err := c.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	schedule, err := saveSchedule(requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, schedule)
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
