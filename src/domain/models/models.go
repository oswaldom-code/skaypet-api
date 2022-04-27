package models

import (
	"time"

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Gender      string    `json:"gender"`
	DateOfBirth time.Time `json:"dateOfBirtth"`
	Specie      string    `json:"specie"`
	Status      bool      `json:"status"`
}

type Specie struct {
	Total  int64  `json:"total"`
	Specie string `json:"specie"`
}

type PetGeneralStatistics struct {
	TotalPets                   int64    `json:"totalPets"`
	AverageAgeYears             int      `json:"averageAgeYears"`
	AverageAgeMonths            int      `json:"averageAgeMonths"`
	NameOfMostNumerousSpecies   string   `json:"mostNumerousSpecies"`
	NumberOfMostNumerousSpecies int      `json:"numberOfMostNumerousSpecies"`
	AgeStdDesviation            float64  `json:"ageStdDesviation"`
	Species                     []Specie `json:"species"`
}

type PetsStatisticsBySpecie struct {
	Specie           string  `json:"specie"`
	TotalPets        int64   `json:"totalPetsBySpecie"`
	AverageAgeYears  int     `json:"averageAgeYears"`
	AverageAgeMonths int     `json:"averageAgeMonths"`
	AgeStdDesviation float64 `json:"ageStdDesviation"`
}

type PetSpecieStatistics struct {
	TotalPets int64  `json:"totalPets"`
	Specie    string `json:"specie"`
}
