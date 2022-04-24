package dto

type MessageResponse struct {
	Status  string `json:"status"`
	Message string `json:"mensaje"`
}

type PetsResponse struct {
	Status string   `json:"status"`
	Pets   []PetDTO `json:"mascotas"`
}

type PetResponse struct {
	Status string `json:"status"`
	Pet    PetDTO `json:"mascota"`
}

type SpecieDTO struct {
	Name  string `json:"nombre"`
	Total int64  `json:"cantidad"`
}

type PetsGeneralStatisticsResponse struct {
	Status             string      `json:"status"`
	TotalPets          int64       `json:"totalMascotas"`
	AverageAgeYears    int         `json:"edadPromedioAnios"`
	AverageAgeMonth    int         `json:"edadPromedioMeses"`
	MostNumerousSpecie string      `json:"specieMasNumerosa"`
	Species            []SpecieDTO `json:"especies"`
}
