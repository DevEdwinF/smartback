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
	sourceCollaborators, err := GetAllColab()
	if err != nil {
		return err
	}

	destinationCollaborators, err := GetAllCollaborators()
	if err != nil {
		return err
	}

	err = syncCollaborators(sourceCollaborators, destinationCollaborators)
	if err != nil {
		return err
	}

	return nil
}

func syncCollaborators(sourceCollaborators []models.NmContr, destinationCollaborators []entity.Collaborators) error {
	for _, sourceCollaborator := range sourceCollaborators {
		found := false

		for i, destinationCollaborator := range destinationCollaborators {
			if sourceCollaborator.Document == destinationCollaborator.Document {
				found = true

				// Actualizar el colaborador en la base de datos destino
				destinationCollaborators[i].FName = sourceCollaborator.FName
				destinationCollaborators[i].LName = sourceCollaborator.LName
				destinationCollaborators[i].Position = sourceCollaborator.Position
				destinationCollaborators[i].Email = sourceCollaborator.EMail
				destinationCollaborators[i].Bmail = sourceCollaborator.BMail
				destinationCollaborators[i].State = sourceCollaborator.State
				destinationCollaborators[i].Leader = sourceCollaborator.FnLeader + " " + sourceCollaborator.LnLeader
				destinationCollaborators[i].LeaderDocument = strconv.Itoa(sourceCollaborator.LeaderDocument)
				destinationCollaborators[i].Subprocess = sourceCollaborator.Subprocess
				destinationCollaborators[i].Headquarters = sourceCollaborator.Headquarters
				destinationCollaborators[i].CreatedAt = time.Now()

				// Actualizar el colaborador en la base de datos destino
				err := UpdateCollaboratorInDestinationDB(destinationCollaborators[i])
				if err != nil {
					return err
				}
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
				LeaderDocument: leaderDocumentStr,
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

func UpdateCollaboratorInDestinationDB(collaborator entity.Collaborators) error {
	err := config.DB.Save(&collaborator).Error
	if err != nil {
		return err
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
