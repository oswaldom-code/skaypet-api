package services

import (
	"math"
	"time"

	"github.com/oswaldom-code/skaypet-api/src/domain/models"
)

func MonthsToYearsMonths(months int) (int, int) {
	years := int(months / 12)
	months = int(months % 12)
	return years, months
}

func DateOfBirthToYearsMonthsDays(dateOfBirth time.Time) (int, int, int) {
	now := time.Now()
	years := now.Year() - dateOfBirth.Year()
	months := int(now.Month()) - int(dateOfBirth.Month())
	days := now.Day() - dateOfBirth.Day()
	// negative values correction in individual cases
	if months < 0 {
		years--
		months = 12 + months
	}
	if days < 0 {
		months--
		days = 30 + days
	}
	return years, months, days
}

func getTotalYearsMonthDaysHelper(pets []models.Pet) (int, int, int, int64) {
	var totalYears, totalMonths, totalDays int
	var TotalSample int64
	// Iterate over all pets to calculate average age in weeks
	for _, pet := range pets {
		years, months, days := DateOfBirthToYearsMonthsDays(pet.DateOfBirth)
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
		TotalSample++
	}
	return totalYears, totalMonths, totalDays, TotalSample
}

// calculateAverageAge return the average age of the pets, return (years, months)
func calculateAverageAge(pets []models.Pet) (int, int, int64) {
	totalYears, totalMonths, totalDays, TotalSample := getTotalYearsMonthDaysHelper(pets)
	// calculate average age
	averageAgeInMonts := int64(totalYears*12+totalMonths+totalDays/30) / TotalSample
	years := int(averageAgeInMonts / 12)
	months := int(averageAgeInMonts % 12)
	return years, months, TotalSample
}

/*
StandardDeviationHelper returns the standard deviation of the ages of
the pets expressed in months

Standard deviation formula:
	S = sqrt(sum((x - x_bar)^2) / n)
	x_bar = (sum(x) / n)
	x = age of pet
	n = total of pets

	variables:
		x = ageExpressedInMonths
		n = TotalSample
		x_bar = averAgeExpressedInMonths
*/

func StandardDeviationInMonths(pets []models.Pet) float64 {
	var averAgeExpressedInMonths, TotalSample, totalDeviation int64
	totalYears, totalMonths, totalDays, TotalSample := getTotalYearsMonthDaysHelper(pets)
	averAgeExpressedInMonths = int64(totalYears*12+totalMonths+totalDays/30) / TotalSample
	for _, pet := range pets {
		years, months, days := DateOfBirthToYearsMonthsDays(pet.DateOfBirth)
		ageExpressedInMonths := int64(years*12 + months + days/30)
		totalDeviation += (ageExpressedInMonths - averAgeExpressedInMonths) * (ageExpressedInMonths - averAgeExpressedInMonths)
	}
	S := math.Sqrt(float64(totalDeviation / TotalSample))
	return S
}

func GetPetStatistics(pets []models.Pet) models.PetGeneralStatistics {
	years, months, TotalSample := calculateAverageAge(pets)
	stdDesviation := StandardDeviationInMonths(pets)

	petGeneralStatistics := models.PetGeneralStatistics{
		TotalPets:        TotalSample,
		AverageAgeYears:  years,
		AverageAgeMonths: months,
		AgeStdDesviation: stdDesviation,
	}
	return petGeneralStatistics
}

func GetMostNumerousSpecieHelper(species []models.Specie) string {
	var maxPets int64
	var mostNumerousSpecie string
	for _, specie := range species {
		if specie.Total > maxPets {
			maxPets = specie.Total
			mostNumerousSpecie = specie.Specie
		}
	}
	return mostNumerousSpecie
}
