package controllers

import (
	"net/http"
	"time"

	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/labstack/echo/v4"
)

func CountAttendancesAll(c echo.Context) error {
	var count int64
	if err := config.DB.Table("attendances").Count(&count).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]int64{"count": count})
}

func CountAttendanceDay(c echo.Context) error {
	var count int64
	today := time.Now().Format("2006-01-02")
	if err := config.DB.Table("attendances").Where("DATE(created_at) = ?", today).Count(&count).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]int64{"count": count})
}
