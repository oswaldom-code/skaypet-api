package repository

import (
	"errors"

	"github.com/oswaldom-code/skaypet-api/pkg/log"
	"github.com/oswaldom-code/skaypet-api/src/domain/models"
)

var (
	SPECIES_AND_COUNT = "SELECT specie, count(*) AS total FROM pets GROUP BY specie HAVING COUNT(*)>1"
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
	pet := models.Pet{}
	db := s.db.First(&pet, id)
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
	pets := []models.Pet{}
	db := s.db.Order("id").Find(&pets)
	if db.Error != nil {
		log.ErrorWithFields("Error getting pets: ", log.Fields{"error": db.Error})
		return pets, db.Error
	}
	return pets, nil
}

func (s *store) existsPet(id int64) bool {
	db := s.db.First(&models.Pet{}, id)
	return db.RowsAffected > 0
}

func (s *store) UpdatePet(id int64, pet models.Pet) (models.Pet, error) {
	if !s.existsPet(id) {
		log.ErrorWithFields("Error updating pet: ", log.Fields{
			"error": "Pet not found",
			"pet":   pet,
		})
		return pet, errors.New("Pet not found")
	}
	db := s.db.Table("pets").Where("id = ?", id).Omit("id").Save(&pet)
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

func (s *store) CountPets() (int64, error) {
	var count int64
	db := s.db.Model(&models.Pet{}).Count(&count)
	if db.Error != nil {
		log.ErrorWithFields("Error counting pets: ", log.Fields{"error": db.Error})
		return 0, db.Error
	}
	return count, nil
}

func (s *store) QuantifySpecies() ([]models.Specie, error) {
	species := []models.Specie{}
	db := s.db.Raw(SPECIES_AND_COUNT).Scan(&species)
	if db.Error != nil {
		log.ErrorWithFields("Error counting pets: ", log.Fields{"error": db.Error})
		return []models.Specie{}, db.Error
	}
	return species, nil
}

func (s *store) GetPetsBySpecie(specie string) ([]models.Pet, error) {
	pets := []models.Pet{}
	db := s.db.Where("specie = ?", specie).Find(&pets)
	if db.Error != nil {
		log.ErrorWithFields("Error getting pets by specie: ", log.Fields{
			"error":  db.Error,
			"specie": specie,
		})
		return pets, db.Error
	}
	if db.RowsAffected == 0 {
		return pets, errors.New("Specie not found: " + specie)
	}
	return pets, nil
}

func (s *store) GetPetsByGender(gender string) ([]models.Pet, error) {
	pets := []models.Pet{}
	db := s.db.Where("gender = ?", gender).Find(&pets)
	if db.Error != nil {
		log.ErrorWithFields("Error getting pets by gender: ", log.Fields{
			"error":  db.Error,
			"gender": gender,
		})
		return pets, db.Error
	}
	if db.RowsAffected == 0 {
		return pets, errors.New("Gender not found: " + gender)
	}
	return pets, nil
}
