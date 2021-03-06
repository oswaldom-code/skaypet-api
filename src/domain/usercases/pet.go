package usercases

import "github.com/oswaldom-code/skaypet-api/src/domain/models"

type PetService interface {
	CreatePet(pet models.Pet) (models.Pet, error)
	GetPet(id int64) (models.Pet, error)
	GetPets() ([]models.Pet, error)
	UpdatePet(pet models.Pet) (models.Pet, error)
	DeletePet(id int64) error
	GetPetsGeneralStatistics() (models.PetGeneralStatistics, error)
	GetQuantifySpecies() []models.Specie
	GetPetsBySpecie(specie string) ([]models.Pet, error)
	GetPetsByGender(gender string) ([]models.Pet, error)
	GetPetsStatisticsBySpecie(gender string) (models.PetsStatisticsBySpecie, error)
}
