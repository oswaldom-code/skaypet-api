package services

import (
	"strings"

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
		log.ErrorWithFields("Error getting pet: ", log.Fields{
			"error": err,
			"petId": id,
		})
		return models.Pet{}, err
	}
	return getPet, nil
}

func (p *petSevice) GetPets() ([]models.Pet, error) {
	getPets, err := p.repository.GetPets()
	if err != nil {
		log.ErrorWithFields("Error getting pets: ", log.Fields{
			"error": err,
		})
		return []models.Pet{}, err
	}
	return getPets, nil
}

func (p *petSevice) UpdatePet(id int64, pet models.Pet) (models.Pet, error) {
	updatePet, err := p.repository.UpdatePet(id, pet)
	if err != nil {
		log.ErrorWithFields("Error updating pet: ", log.Fields{
			"error": err,
			"petId": id,
			"pet":   pet,
		})
		return models.Pet{}, err
	}
	return updatePet, nil
}

func (p *petSevice) DeletePet(id int64) error {
	if err := p.repository.DeletePet(id); err != nil {
		log.ErrorWithFields("Error deleting pet: ", log.Fields{
			"error": err,
			"petId": id,
		})
		return err
	}
	return nil
}

func (p *petSevice) CountPets() (int64, error) {
	TotalPets, err := p.repository.CountPets()
	if err != nil {
		log.ErrorWithFields("Error counting pets: ", log.Fields{
			"error": err,
		})
		return 0, err
	}
	return TotalPets, nil
}

func (p *petSevice) GetPetsGeneralStatistics() (models.PetGeneralStatistics, error) {
	pets, err := p.repository.GetPets()
	if err != nil {
		log.ErrorWithFields("Error getting pets: ", log.Fields{
			"error": err,
		})
		return models.PetGeneralStatistics{}, err
	}
	petGeneralStatistics := GetPetGeneralStatisticsHelper(pets)
	petGeneralStatistics.Species = p.GetQuantifySpecies()
	petGeneralStatistics.NameOfMostNumerousSpecies = GetMostNumerousSpecieHelper(petGeneralStatistics.Species)
	return petGeneralStatistics, nil
}

func (p *petSevice) GetQuantifySpecies() []models.Specie {
	quantifySpecies, err := p.repository.QuantifySpecies()
	if err != nil {
		log.ErrorWithFields("Error GetQuantifySpecies species: ", log.Fields{
			"error": err,
		})
		return []models.Specie{}
	}
	return quantifySpecies
}

func (p *petSevice) GetPetsBySpecie(specie string) ([]models.Pet, error) {
	pets, err := p.repository.GetPetsBySpecie(strings.ToUpper(specie))
	if err != nil {
		log.ErrorWithFields("Error getting pets by specie: ", log.Fields{
			"error":  err,
			"specie": specie,
		})
		return []models.Pet{}, err
	}
	return pets, nil
}
