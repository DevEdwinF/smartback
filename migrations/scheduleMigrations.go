package migrations

import (
	"net/http"

	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

type ScheduleMigrationService struct{}

func NewScheduleMigrationService() *ScheduleMigrationService {
	return &ScheduleMigrationService{}
}

// func (s *ScheduleMigrationService) MigrationSchedule(schedule models.Schedules) {
// }

func UploadExcelAndAssignSchedules(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error al cargar el archivo")
	}

	excelFile, err := excelize.OpenFile(file.Filename)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error al abrir el archivo Excel")
	}

	rows, err := excelFile.GetRows("horario")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error al leer las filas del archivo Excel")
	}

	for _, row := range rows {
		document := row[0]
		day := row[1]
		arrivalTime := row[2]
		departureTime := row[3]

		var collaborator entity.CollaboratorsDataEntity
		err := config.DB.Table("collaborators").
			Select("id").
			Where("document = ?", document).
			First(&collaborator).Error
		if err != nil {
			continue
		}

		schedule := entity.Schedules{
			Document:         document,
			Day:              day,
			ArrivalTime:      arrivalTime,
			DepartureTime:    departureTime,
			FkCollaboratorId: collaborator.ID,
		}

		err = config.DB.Table("schedules").Create(&schedule).Error
		if err != nil {
			continue
		}
	}

	return c.JSON(http.StatusOK, "Carga masiva completada")
}
