package port

import "github.com/oswaldom-code/skaypet-api/src/domain/models"

type Repository interface {
	// Pet
	CreatePet(pet models.Pet) (models.Pet, error)
	GetPet(id int64) (models.Pet, error)
	GetPets() ([]models.Pet, error)
	UpdatePet(pet models.Pet) (models.Pet, error)
	DeletePet(id int64) error
	CountPets() (int64, error)
	CountPetsBySpecie(specie string) (int64, error)
	QuantifySpecies() ([]models.Specie, error)
	GetPetsBySpecie(string) ([]models.Pet, error)
	GetPetsByGender(string) ([]models.Pet, error)
}
