package dto

import (
	"strings"

	"github.com/oswaldom-code/skaypet-api/src/domain/models"
)

type PetDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"nombre"`
	Age         string `json:"edad"`
	Gender      string `json:"genero"`
	DateOfBirth string `json:"fechaNacimiento"`
	Status      bool   `json:"status"`
	Specie      string `json:"specie"`
}

func PetToPetDTO(pet models.Pet) PetDTO {
	petDTO := PetDTO{
		ID:          pet.ID,
		Name:        pet.Name,
		Age:         AgeInYearsMonthsDays(pet.DateOfBirth),
		Gender:      pet.Gender,
		DateOfBirth: pet.DateOfBirth.Format("2006-01-02"),
		Status:      pet.Status,
		Specie:      pet.Specie,
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
			Specie:      pet.Specie,
		}
		petsDTO = append(petsDTO, petDTO)
	}
	return petsDTO
}

func PetDTOToPet(PetDTO PetDTO) models.Pet {
	pet := models.Pet{
		ID:          PetDTO.ID,
		Name:        strings.ToUpper(PetDTO.Name),
		Gender:      strings.ToUpper(PetDTO.Gender),
		DateOfBirth: TimeStringToTime(PetDTO.DateOfBirth),
		Status:      PetDTO.Status,
		Specie:      strings.ToUpper(PetDTO.Specie),
	}
	return pet
}

func PetsGeneralStatisticsToPetsGeneralStatisticsDTO(petsStatistics models.PetGeneralStatistics) PetsGeneralStatisticsResponse {
	petsGeneralStatisticsResponse := PetsGeneralStatisticsResponse{
		Status:             "success",
		TotalPets:          petsStatistics.TotalPets,
		AverageAgeYears:    petsStatistics.AverageAgeYears,
		AverageAgeMonths:   petsStatistics.AverageAgeMonths,
		MostNumerousSpecie: petsStatistics.NameOfMostNumerousSpecies,
		AgeStdDesviation:   petsStatistics.AgeStdDesviation,
		Species:            SpeciesToSpeciesDTO(petsStatistics.Species),
	}
	return petsGeneralStatisticsResponse
}

func PetsStatisticsBySpecieToPetsStatisticsBySpecieDTO(petsStatistics models.PetsStatisticsBySpecie) PetsStatisticsBySpecieResponse {
	petsGeneralStatisticsResponse := PetsStatisticsBySpecieResponse{
		Status:           "success",
		Specie:           petsStatistics.Specie,
		TotalPets:        petsStatistics.TotalPets,
		AverageAgeYears:  petsStatistics.AverageAgeYears,
		AverageAgeMonths: petsStatistics.AverageAgeMonths,
		AgeStdDesviation: petsStatistics.AgeStdDesviation,
	}
	return petsGeneralStatisticsResponse
}

func SpeciesToSpeciesDTO(species []models.Specie) []SpecieDTO {
	speciesDTO := []SpecieDTO{}
	for _, specie := range species {
		specieDTO := SpecieDTO{
			Name:  specie.Specie,
			Total: specie.Total,
		}
		speciesDTO = append(speciesDTO, specieDTO)
	}
	return speciesDTO
}
