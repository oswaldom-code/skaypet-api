package dto

import (
	"github.com/oswaldom-code/skaypet-api/src/domain/models"
)

type PetDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"nombre"`
	Age         string `json:"edad"`
	Gender      string `json:"genero"`
	DateOfBirth string `json:"fechaNacimiento"`
	Status      string `json:"status"`
	Specie      int64 `json:"specieId"`
}

func PetsToPetsDTO(pets []models.Pet) []PetDTO {
	petsDTO := []PetDTO{}
	for _, pet := range pets {
		petDTO := PetDTO{
			ID:   pet.ID,
			Name: pet.Name,
			Age:      AgeInYearsMonthsDays(pet.DateOfBirth),
			Gender:      pet.Gender,
			DateOfBirth: pet.DateOfBirth.Format("2006-01-02"),
			Status:      pet.Status,
			Specie:	pet.SpecieId,
		}
		petsDTO = append(petsDTO, petDTO)
	}
	return petsDTO
}
