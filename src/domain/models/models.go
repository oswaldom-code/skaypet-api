package models

import (
	"time"

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	ID          int64     
	Name        string    
	Gender      string    
	DateOfBirth time.Time 
	SpecieId    int64     `gorm:"foreignKey:SpecieId;references:ID"`
	Status      string    
}

type Specie struct {
	// omitir parametro ID en respuesta json
	ID   int64  `json:"specie_id"`
	Name string  `json:"specie_nombre"`
}
