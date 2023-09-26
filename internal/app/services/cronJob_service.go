package services

import (
	"log"
	"strconv"
	"time"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"github.com/robfig/cron"
)

func RunCronJob() {
	c := cron.New()

	// c.AddFunc("0 */30 * * * *", func() {
	c.AddFunc("*/5 * * * *", func() {
		err := SyncData()
		if err != nil {
			log.Println("Error al sincronizar datos:", err)
		}
	})
	// c.AddFunc("0 0 22 * * *", func() {
	// 	err := SyncData()
	// 	if err != nil {
	// 		log.Println("Error al sincronizar datos:", err)
	// 	}
	// })

	c.Start()

	select {}
}

func SyncData() error {
	// Obtener los colaboradores de la base de datos fuente
	sourceCollaborators, err := GetAllColab()
	if err != nil {
		return err
	}

	// Obtener los colaboradores de la base de datos de destino
	destinationCollaborators, err := GetAllCollaborators()
	if err != nil {
		return err
	}

	// Realizar la comparación y alimentar la base de datos de destino con nuevos datos
	err = syncCollaborators(sourceCollaborators, destinationCollaborators)
	if err != nil {
		return err
	}

	return nil
}

func syncCollaborators(sourceCollaborators []models.NmContr, destinationCollaborators []entity.Collaborators) error {
	for _, sourceCollaborator := range sourceCollaborators {
		found := false

		for _, destinationCollaborator := range destinationCollaborators {
			if sourceCollaborator.Document == destinationCollaborator.Document {
				found = true
				break
			}
		}

		if !found {
			leaderDocumentStr := strconv.Itoa(sourceCollaborator.LeaderDocument)

			newCollaborator := entity.Collaborators{
				Document:       sourceCollaborator.Document,
				FName:          sourceCollaborator.FName,
				LName:          sourceCollaborator.LName,
				Position:       sourceCollaborator.Position,
				Email:          sourceCollaborator.EMail,
				Bmail:          sourceCollaborator.BMail,
				State:          sourceCollaborator.State,
				Leader:         sourceCollaborator.FnLeader + " " + sourceCollaborator.LnLeader,
				LeaderDocument: leaderDocumentStr, // Asignar la versión en string de LeaderDocument
				Subprocess:     sourceCollaborator.Subprocess,
				Headquarters:   sourceCollaborator.Headquarters,
				CreatedAt:      time.Now(),
			}

			err := AddCollaboratorToDestinationDB(newCollaborator)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func AddCollaboratorToDestinationDB(collaborator entity.Collaborators) error {
	err := config.DB.Create(&collaborator).Error
	if err != nil {
		return err
	}

	return nil
}
