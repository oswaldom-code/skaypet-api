package repository

import "github.com/oswaldom-code/skaypet-api/src/domain/models"

type PetRepository interface {
	GetPets() ([]models.Pet, error)
}
