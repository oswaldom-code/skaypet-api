package repository

import (
	"github.com/oswaldom-code/skaypet-api/src/domain/models"
)

func (s *store) GetPets() ([]models.Pet, error) {
	var pets []models.Pet
	db := s.db.Table("pet").Joins("inner join specie on pet.specie_id = specie.id").Order("id").Find(&pets)
	if db.Error != nil {
		//TODO log message
		return pets, db.Error
	}
	return pets, nil
}
