package services

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"github.com/DevEdwinF/smartback.git/internal/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AttendanceService struct{}

func NewAttendanceService() *AttendanceService {
	return &AttendanceService{}
}

func (s *AttendanceService) GetUnregisteredForDay(attendance entity.AttendanceEntity, day time.Time) ([]entity.Collaborators, error) {
	allCollaborators, err := ValidateCollaborator()
	if err != nil {
		return nil, err
	}

	attendanceForDay, err := GetAttendanceForDay(day)
	if err != nil {
		return nil, err
	}

	attendanceMap := make(map[string]bool)
	for _, attendance := range attendanceForDay {
		attendanceMap[attendance.Document] = true
	}

	unregisteredCollaborators := []entity.Collaborators{}
	for _, collaborator := range allCollaborators {
		if _, ok := attendanceMap[collaborator.Document]; !ok {
			unregisteredCollaborators = append(unregisteredCollaborators, collaborator)
		}
	}

	return unregisteredCollaborators, nil
}

func GetAttendanceForDay(day time.Time) ([]entity.AttendanceEntity, error) {
	attendanceForDay := []entity.AttendanceEntity{}

	if err := config.DB.Where("DATE(created_at) = ?", day.Format("2006-01-02")).Find(&attendanceForDay).Error; err != nil {
		return nil, err
	}

	return attendanceForDay, nil
}

func (s *AttendanceService) RegisterAttendance(attendance entity.AttendanceEntity) error {
	collaborator, err := GetCollaborator(attendance.Document)
	if err != nil {
		return err
	}

	if collaborator == nil {
		return errors.New("Colaborador no encontrado")
	}

	loc, err := time.LoadLocation("America/Bogota")

	if err != nil {
		return err
	}
	timeNow := time.Now().In(loc)

	var schedule models.Schedules
	err = config.DB.Model(&schedule).Where("fk_collaborator_id = ? AND day = ?", collaborator.Id, timeNow.Format("Monday")).First(&schedule).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return errors.New("Horario no encontrado para el colaborador en este día")
	}

	var arrivalScheduled time.Time
	if schedule.ArrivalTime != "" {
		temp, err := time.Parse("15:04:05", schedule.ArrivalTime)
		if err != nil {
			return err
		}
		arrivalScheduled = time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), temp.Hour(), temp.Minute(), temp.Second(), temp.Nanosecond(), timeNow.Location())
	}

	late := false

	if !arrivalScheduled.IsZero() && timeNow.After(arrivalScheduled.Add(6*time.Minute)) {
		late = true
	}

	var departureScheduled time.Time

	if schedule.DepartureTime != "" {
		temp, err := time.Parse("15:04:05", schedule.DepartureTime)
		if err != nil {
			return err
		}
		departureScheduled = time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), temp.Hour(), temp.Minute(), temp.Second(), temp.Nanosecond(), timeNow.Location())
	}

	earlyDeparture := false

	if !departureScheduled.IsZero() && timeNow.After(departureScheduled) {
		earlyDeparture = true
	}

	var validateAttendance models.Attendance
	err = config.DB.Model(&validateAttendance).
		Where("fk_collaborator_id = ? AND date(created_at) = ?", collaborator.Id, timeNow.Format("2006-01-02")).
		First(&validateAttendance).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	/* folderPath := "/app/attendance_photos" */
	folderPath := "attendance_photos"
	err = os.MkdirAll(folderPath, 0755)
	if err != nil {
		log.Println("Error creating directory:", err)
		return err
	}

	imagenCodificadaEnBase64 := attendance.Photo

	marca := ";base64,"
	indice := strings.Index(imagenCodificadaEnBase64, marca)
	if indice != -1 {
		imagenCodificadaEnBase64 = imagenCodificadaEnBase64[indice+len(marca):]
	}

	decodificado, err := base64.StdEncoding.DecodeString(imagenCodificadaEnBase64)
	if err != nil {
		log.Println("Error decoding base64 image:", err)
		return err
	}

	imagen, _, err := image.Decode(bytes.NewReader(decodificado))
	if err != nil {
		log.Println("Error saving image:", err)
		return err
	}

	photoName := fmt.Sprintf("%s%d.png", "1150856537", time.Now().Unix())

	archivo, err := os.Create(fmt.Sprintf("%s/%s", folderPath, photoName))
	if err != nil {
		log.Println("Error al crear el archivos:", err)
		return err
	}
	err = png.Encode(archivo, imagen)
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	err = archivo.Close()
	if err != nil {
		log.Println("Error al cerrar el archivo:", err)
		return err
	}

	attendance.Photo = photoName

	switch attendance.State {
	case "arrival":
		if validateAttendance.ID != 0 || validateAttendance.Arrival.Valid {
			return errors.New("Ya se ha registrado la entrada")
		}

		modelsAttendance := models.Attendance{
			FkCollaboratorID: collaborator.Id,
			PhotoArrival:     attendance.Photo,
			Location:         attendance.Location,
			Arrival:          sql.NullString{String: timeNow.Format("15:04:05"), Valid: true},
			Late:             late,
			EarlyDeparture:   earlyDeparture,
			CreatedAt:        timeNow,
		}
		err = config.DB.Create(&modelsAttendance).Error
		if err != nil {
			return err
		}

		return nil

	case "departure":
		if validateAttendance.ID == 0 {
			return errors.New("Debe registrar la entrada primero")
		}
		if validateAttendance.Departure.Valid {
			return errors.New("Ya se ha registrado la salida")
		}

		validateAttendance.Departure = sql.NullString{String: timeNow.Format("15:04:05"), Valid: true}
		validateAttendance.PhotoDeparture = attendance.Photo
		err = config.DB.Updates(&validateAttendance).Error
		if err != nil {
			return err
		}

		return nil
	}

	return errors.New("Estado inválido")
}

func (service *AttendanceService) GetAttendancePage(filter entity.AttendanceFilter) (entity.Pagination, error) {
	offset := (filter.Page - 1) * filter.Limit
	var count int64

	attendance := []entity.UserAttendanceData{}
	var where string
	utils.BuildFilters("document", filter.Document, "OR", &where)
	utils.BuildFilters("f_name", filter.FName, "OR", &where)
	utils.BuildFilters("l_name", filter.LName, "OR", &where)
	utils.BuildFilters("bmail", filter.Bmail, "OR", &where)
	utils.BuildFilters("email", filter.Email, "OR", &where)
	utils.BuildFilters("arrival", filter.Arrival, "OR", &where)
	utils.BuildFilters("departure", filter.Departure, "OR", &where)
	utils.BuildFilters("leader", filter.Leader, "OR", &where)
	utils.BuildFilters("position", filter.Position, "OR", &where)
	utils.BuildFilters("headqarters", filter.Headquarters, "OR", &where)
	utils.BuildFilters("subprocess", filter.Subprocess, "OR", &where)
	utils.BuildFilters("location", filter.Location, "OR", &where)

	err := config.DB.Table("attendances a").
		Select("c.f_name, c.l_name, c.email, c.document, a.*").
		Joins("INNER JOIN collaborators c on c.id = a.fk_collaborator_id").
		Where(where).
		Count(&count).
		Offset(offset).Limit(filter.Limit).
		Scan(&attendance).Error
	if err != nil {
		return entity.Pagination{}, err
	}

	folderPath := "attendance_photos"

	for i := range attendance {
		photoArrivalName := attendance[i].PhotoArrival
		photoArrivalPath := filepath.Join(folderPath, photoArrivalName)

		imageArrivalData, err := ioutil.ReadFile(photoArrivalPath)
		if err == nil {
			base64ImageArrival := base64.StdEncoding.EncodeToString(imageArrivalData)
			attendance[i].PhotoArrival = base64ImageArrival
		}

		photoDepartureName := attendance[i].PhotoDeparture
		photoDeparturePath := filepath.Join(folderPath, photoDepartureName)

		imageDepartureData, err := ioutil.ReadFile(photoDeparturePath)
		if err == nil {
			base64ImageDeparture := base64.StdEncoding.EncodeToString(imageDepartureData)
			attendance[i].PhotoDeparture = base64ImageDeparture
		}
	}

	return entity.Pagination{
		Page:      filter.Page,
		Limit:     filter.Limit,
		TotalPage: int(math.Ceil(float64(count) / float64(filter.Limit))),
		TotalRows: count,
		Rows:      attendance,
	}, nil
}

func (service *AttendanceService) GetAttendanceForLeaderPage(filter entity.AttendanceFilter) (entity.Pagination, error) {
	offset := (filter.Page - 1) * filter.Limit
	var count int64

	attendance := []entity.UserAttendanceData{}

	var where string
	utils.BuildFilters("document", filter.Document, "OR", &where)
	utils.BuildFilters("f_name", filter.FName, "OR", &where)
	utils.BuildFilters("l_name", filter.LName, "OR", &where)
	utils.BuildFilters("b_mail", filter.Bmail, "OR", &where)
	utils.BuildFilters("email", filter.Email, "OR", &where)
	utils.BuildFilters("position", filter.Position, "OR", &where)
	utils.BuildFilters("leader", filter.Leader, "OR", &where)
	utils.BuildFilters("subprocess", filter.Subprocess, "OR", &where)
	utils.BuildFilters("position", filter.Position, "OR", &where)
	utils.BuildFilters("leader_document", filter.LeaderDocument, "AND", &where)

	err := config.DB.Table("attendances a").
		Select("c.f_name, c.l_name, c.email, c.leader, c.document, a.*").
		Joins("INNER JOIN collaborators c ON c.id = a.fk_collaborator_id").
		Where(where).
		Order("created_at DESC").
		Count(&count).
		Offset(offset).
		Limit(filter.Limit).
		Scan(&attendance).Error

	if err != nil {
		return entity.Pagination{}, err
	}

	folderPath := "attendance_photos"

	for i := range attendance {
		// Process PhotoArrival
		photoArrivalName := attendance[i].PhotoArrival
		photoArrivalPath := filepath.Join(folderPath, photoArrivalName)

		imageArrivalData, err := ioutil.ReadFile(photoArrivalPath)
		if err == nil {
			base64ImageArrival := base64.StdEncoding.EncodeToString(imageArrivalData)
			attendance[i].PhotoArrival = base64ImageArrival
		}

		// Process PhotoDeparture
		photoDepartureName := attendance[i].PhotoDeparture
		photoDeparturePath := filepath.Join(folderPath, photoDepartureName)

		imageDepartureData, err := ioutil.ReadFile(photoDeparturePath)
		if err == nil {
			base64ImageDeparture := base64.StdEncoding.EncodeToString(imageDepartureData)
			attendance[i].PhotoDeparture = base64ImageDeparture
		}
	}

	return entity.Pagination{
		Page:      filter.Page,
		Limit:     filter.Limit,
		TotalPage: int(math.Ceil(float64(count) / float64(filter.Limit))),
		TotalRows: count,
		Rows:      attendance,
	}, nil
}

func (service *AttendanceService) GetAllAttendanceForToLate() ([]entity.UserAttendanceData, error) {
	attendance := []entity.UserAttendanceData{}
	err := config.DB.Table("collaborators c").
		Select("c.f_name, c.l_name, c.email, c.leader, c.document, a.*").
		Joins("INNER JOIN attendances a ON c.id = a.fk_collaborator_id").
		Where("a.late = TRUE").
		Where("EXISTS (SELECT 1 FROM attendances WHERE fk_collaborator_id = c.id AND late = TRUE HAVING COUNT(*) > 2)").
		Find(&attendance).Error
	if err != nil {
		return nil, err
	}

	return attendance, nil
}

// func (service *AttendanceService) GetAttendanceForLeaderToLate(leaderFullName string) ([]entity.UserAttendanceData, error) {
// 	attendance := []entity.UserAttendanceData{}
// 	err := config.DB.Table("collaborators c").
// 		Select("c.f_name, c.l_name, c.email, c.leader, c.document, a.*").
// 		Joins("INNER JOIN users u ON CONCAT(u.f_name, ' ', u.l_name) = c.leader").
// 		Joins("INNER JOIN attendances a ON c.id = a.fk_collaborator_id").
// 		Where("c.leader = ?", leaderFullName).
// 		Where("EXISTS (SELECT 1 FROM attendances WHERE fk_collaborator_id = c.id AND late = TRUE HAVING COUNT(*) > 2)").
// 		Find(&attendance).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	folderPath := "attendance_photos"

// 	for i := range attendance {
// 		photoName := attendance[i].Photo
// 		imagePath := filepath.Join(folderPath, photoName)

// 		imageData, err := ioutil.ReadFile(imagePath)
// 		if err != nil {
// 			return nil, err
// 		}

// 		base64Image := base64.StdEncoding.EncodeToString(imageData)

// 		attendance[i].Photo = base64Image
// 	}

// 	return attendance, nil
// }

func (service *AttendanceService) GetAttendanceForLeaderToLate(leaderDocument string) ([]entity.UserAttendanceData, error) {
	attendance := []entity.UserAttendanceData{}
	err := config.DB.Table("collaborators c").
		Select("c.f_name, c.l_name, c.email, c.leader, c.document, a.*").
		Joins("INNER JOIN users u ON u.document = c.leader_document").
		Joins("INNER JOIN attendances a ON c.id = a.fk_collaborator_id").
		Where("u.document = ?", leaderDocument).
		Where("a.late = TRUE").
		Where("EXISTS (SELECT 1 FROM attendances WHERE fk_collaborator_id = c.id AND late = TRUE HAVING COUNT(*) > 2)").
		Find(&attendance).Error
	if err != nil {
		return nil, err
	}

	return attendance, nil
}

func (service *AttendanceService) SaveTranslatedService(translatedEntity entity.Translatedcollaborators) error {
	var collaborator models.Collaborators
	err := config.DB.Model(&collaborator).Where("document = ?", translatedEntity.Document).First(&collaborator).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return errors.New("Colaborador no encontrado")
	}

	loc, err := time.LoadLocation("America/Bogota")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	currentTime := time.Now().In(loc)

	newTranslatedCollaborator := models.Translatedcollaborators{
		FkCollaboratorId: collaborator.Id,
		CreatedAt:        currentTime,
	}

	err = config.DB.Create(&newTranslatedCollaborator).Error
	if err != nil {
		return err
	}

	return nil
}
