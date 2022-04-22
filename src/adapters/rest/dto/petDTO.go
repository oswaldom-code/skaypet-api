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
	Status      bool   `json:"status"`
	Specie      int64  `json:"specieId"`
}

func PetToPetDTO(pet models.Pet) PetDTO {
	petDTO := PetDTO{
		ID:          pet.ID,
		Name:        pet.Name,
		Age:         AgeInYearsMonthsDays(pet.DateOfBirth),
		Gender:      pet.Gender,
		DateOfBirth: pet.DateOfBirth.Format("2006-01-02"),
		Status:      pet.Status,
		Specie:      pet.SpecieId,
	}
	return petDTO
}

func PetsToPetsDTO(pets []models.Pet) []PetDTO {
	petsDTO := []PetDTO{}
	for _, pet := range pets {
		petDTO := PetDTO{
			ID:          pet.ID,
			Name:        pet.Name,
			Age:         AgeInYearsMonthsDays(pet.DateOfBirth),
			Gender:      pet.Gender,
			DateOfBirth: pet.DateOfBirth.Format("2006-01-02"),
			Status:      pet.Status,
			Specie:      pet.SpecieId,
		}
		petsDTO = append(petsDTO, petDTO)
	}
	return petsDTO
}

func PetDTOToPet(PetDTO PetDTO) models.Pet {
	pet := models.Pet{
		ID:   PetDTO.ID,
		Name: PetDTO.Name,
		//Age:         AgeInYearsMonthsDays(PetDTO.DateOfBirth),
		Gender:      PetDTO.Gender,
		DateOfBirth: TimeStringToTime(PetDTO.DateOfBirth),
		Status:      PetDTO.Status,
		SpecieId:    PetDTO.Specie,
	}
	return pet
}
