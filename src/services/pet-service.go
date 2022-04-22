package services

import (
	"github.com/oswaldom-code/skaypet-api/pkg/log"
	"github.com/oswaldom-code/skaypet-api/src/adapters/persistence/repository"
	"github.com/oswaldom-code/skaypet-api/src/domain/models"
	"github.com/oswaldom-code/skaypet-api/src/domain/usercases"
	"github.com/oswaldom-code/skaypet-api/src/services/port"
)

type petSevice struct {
	repository port.Repository
}

func NewPetSevice() usercases.PetService {
	return &petSevice{repository: repository.GetPersistence()}
}

func (p *petSevice) CreatePet(pet models.Pet) (models.Pet, error) {
	createPet, err := p.repository.CreatePet(pet)
	if err != nil {
		log.ErrorWithFields("Error creating pet: ", log.Fields{
			"error": err,
			"pet":   pet,
		})

		return models.Pet{}, err
	}
	return createPet, nil
}

func (p *petSevice) GetPet(id int64) (models.Pet, error) {
	getPet, err := p.repository.GetPet(id)
	if err != nil {
		return models.Pet{}, err
	}
	return getPet, nil
}

func (p *petSevice) GetPets() ([]models.Pet, error) {
	getPets, err := p.repository.GetPets()
	if err != nil {
		return []models.Pet{}, err
	}
	return getPets, nil
}

func (p *petSevice) UpdatePet(id int64, pet models.Pet) (models.Pet, error) {
	updatePet, err := p.repository.UpdatePet(id, pet)
	if err != nil {
		return models.Pet{}, err
	}
	return updatePet, nil
}

func (p *petSevice) DeletePet(id int64) error {
	if err := p.repository.DeletePet(id); err != nil {
		return err
	}
	return nil
}

func (p *petSevice) CountPets() (int64, error) {
	TotalPets, err := p.repository.CountPets()
	if err != nil {
		return 0, err
	}
	return TotalPets, nil
}
