package usercases

import "github.com/oswaldom-code/skaypet-api/src/domain/models"

type PetService interface {
	GetPets() ([]models.Pet, error)
}
