package services

import (
	"github.com/oswaldom-code/skaypet-api/src/adapters/persistence/repository"
	"github.com/oswaldom-code/skaypet-api/src/domain/models"
	"github.com/oswaldom-code/skaypet-api/src/domain/usercases"
)


type petSevice struct {
	repository repository.Repository
}


func NewPetSevice() usercases.PetService {
	return &petSevice{repository: repository.GetPersistence()}
}

func (p *petSevice) GetPets() ([]models.Pet, error) {
	getPets, err := p.repository.GetPets()
	if err != nil {
		return []models.Pet{}, err
	}
	return getPets, nil
}
