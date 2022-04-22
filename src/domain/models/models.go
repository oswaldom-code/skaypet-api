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
	SpecieId    int64     `gorm:"foreignKey:SpecieId;references:ID"`
	Status      bool      `json:"status"`
}

type Specie struct {
	ID   int64  `json:"specie_id"`
	Name string `json:"specie_nombre"`
}
