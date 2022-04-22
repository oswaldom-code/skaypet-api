package repository

import (
	"github.com/oswaldom-code/skaypet-api/pkg/log"
	"github.com/oswaldom-code/skaypet-api/src/domain/models"
)

func (s *store) CreatePet(pet models.Pet) (models.Pet, error) {
	db := s.db.Create(&pet)
	if db.Error != nil {
		log.ErrorWithFields("Error creating pet: ", log.Fields{
			"error": db.Error,
			"pet":   pet,
		})
		return pet, db.Error
	}
	return pet, nil
}

func (s *store) GetPet(id int64) (models.Pet, error) {
	var pet models.Pet
	db := s.db.Table("pet").Joins("inner join specie on pet.specie_id = specie.id").Where("pet.id = ?", id).First(&pet)
	if db.Error != nil {
		log.ErrorWithFields("Error getting pet: ", log.Fields{
			"error": db.Error,
			"pet":   pet,
		})
		return pet, db.Error
	}
	return pet, nil
}

func (s *store) GetPets() ([]models.Pet, error) {
	var pets []models.Pet
	db := s.db.Table("pet").Joins("inner join specie on pet.specie_id = specie.id").Order("id").Find(&pets)
	if db.Error != nil {
		log.ErrorWithFields("Error getting pets: ", log.Fields{"error": db.Error})
		return pets, db.Error
	}
	return pets, nil
}

func (s *store) UpdatePet(pet models.Pet) (models.Pet, error) {
	db := s.db.Save(&pet)
	if db.Error != nil {
		log.ErrorWithFields("Error updating pet: ", log.Fields{
			"error": db.Error,
			"pet":   pet,
		})
		return pet, db.Error
	}
	return pet, nil
}

func (s *store) DeletePet(id int64) error {
	db := s.db.Delete(&models.Pet{}, id)
	if db.Error != nil {
		log.ErrorWithFields("Error deleting pet: ", log.Fields{
			"error": db.Error,
			"petId": id,
		})
		return db.Error
	}
	return nil
}
