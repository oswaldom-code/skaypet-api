package services

import (
	"time"

	"github.com/oswaldom-code/skaypet-api/src/domain/models"
)

// calculateAverageAge return the average age of the pets, return (years, months)
func calculateAverageAge(pets []models.Pet) (int, int) {
	now := time.Now()
	var totalYears, totalMonths, totalDays int
	var totalPets int64
	// Iterate over all pets to calculate average age in weeks
	for _, pet := range pets {
		years := now.Year() - pet.DateOfBirth.Year()
		months := int(now.Month()) - int(pet.DateOfBirth.Month())
		days := now.Day() - pet.DateOfBirth.Day()
		// negative values correction in individual cases
		if months < 0 {
			years--
			months = 12 + months
		}
		if days < 0 {
			months--
			days = 30 + days
		}
		// acumulete values of years and months
		totalYears += years
		totalMonths += months
		if totalMonths > 12 {
			totalYears += totalMonths / 12
			totalMonths = totalMonths % 12
		}
		totalDays += days
		if totalDays > 30 {
			totalMonths += totalDays / 30
			totalDays = totalDays % 30
		}
		totalPets++
	}
	// calculate average age
	averageAge := int64(totalYears*12+totalMonths+totalDays/30) / totalPets
	years := int(averageAge / 12)
	months := int(averageAge % 12)
	return years, months
}

func GetPetGeneralStatisticsHelper(pets []models.Pet) models.PetGeneralStatistics {
	years, months := calculateAverageAge(pets)
	petGeneralStatistics := models.PetGeneralStatistics{
		TotalPets:       int64(len(pets)),
		AverageAgeYears: years,
		AverageAgeMonth: months,
	}
	return petGeneralStatistics
}
