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
	AverageAgeYears             int      `json:"averageAge"`
	AverageAgeMonth             int      `json:"averageAgeMonth"`
	NameOfMostNumerousSpecies   string   `json:"mostNumerousSpecies"`
	NumberOfMostNumerousSpecies int      `json:"numberOfMostNumerousSpecies"`
	Species                     []Specie `json:"species"`
}

type PetSpecieStatistics struct {
	TotalPets int64  `json:"totalPets"`
	Specie    string `json:"specie"`
}
