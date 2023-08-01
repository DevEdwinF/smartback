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

//testing

func syncCollaborators(sourceCollaborators []models.NmContr, destinationCollaborators []entity.Collaborators) error {
	// Recorre los colaboradores de la base de datos fuente
	for _, sourceCollaborator := range sourceCollaborators {
		found := false

		// Busca el colaborador en la base de datos de destino
		for _, destinationCollaborator := range destinationCollaborators {
			if sourceCollaborator.Document == strconv.Itoa(destinationCollaborator.Document) {
				// El colaborador ya existe en la base de datos de destino, no es necesario agregarlo
				found = true
				break
			}
		}

		if !found {
			// El colaborador no existe en la base de datos de destino, agrega el nuevo colaborador

			// Convertir el campo 'Document' de string a int
			documentInt, err := strconv.Atoi(sourceCollaborator.Document)
			if err != nil {
				return err // Manejar el error si el valor no es un número válido
			}

			newCollaborator := entity.Collaborators{
				Document: documentInt,
				FName:    sourceCollaborator.FName,
				LName:    sourceCollaborator.LName,
				Position: sourceCollaborator.Position,
				Leader:   sourceCollaborator.FnLeader + " " + sourceCollaborator.LnLeader,
				CreateAt: time.Now(),
			}

			err = AddCollaboratorToDestinationDB(newCollaborator) // Use the existing 'err' variable
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
