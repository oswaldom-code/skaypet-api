package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oswaldom-code/skaypet-api/src/adapters/rest/dto"
	"github.com/oswaldom-code/skaypet-api/src/services"
)

func GetPets(c *gin.Context) {
	petSevice := services.NewPetSevice()
	pets, err := petSevice.GetPets()
	if err != nil {
		response := dto.MessageResponse{
			Status:  "error",
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := dto.PetsResponse{
		Status: "success",
		Pets:   dto.PetsToPetsDTO(pets),
	}
	c.JSON(http.StatusOK, response)
}