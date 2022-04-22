package dto

type MessageResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type PetsResponse struct {
	Status string `json:"status"`
	Pets   []PetDTO  `json:"pets"`
}