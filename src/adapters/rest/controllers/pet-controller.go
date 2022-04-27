package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oswaldom-code/skaypet-api/src/adapters/rest/dto"
	"github.com/oswaldom-code/skaypet-api/src/services"
)

func Pong(c *gin.Context) {
	response := dto.MessageResponse{
		Status:  "success",
		Message: "pong",
	}
	c.JSON(http.StatusOK, response)
}

func CreatePet(c *gin.Context) {
	petSevice := services.NewPetSevice()
	var petDTO dto.PetDTO
	if err := c.ShouldBindJSON(&petDTO); err != nil {
		response := dto.MessageResponse{
			Status:  "error",
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	pet, err := petSevice.CreatePet(dto.PetDTOToPet(petDTO))
	if err != nil {
		response := dto.MessageResponse{
			Status:  "error",
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := dto.PetResponse{
		Status: "success",
		Pet:    dto.PetToPetDTO(pet),
	}
	c.JSON(http.StatusCreated, response)
}

func GetPet(c *gin.Context) {
	petSevice := services.NewPetSevice()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response := dto.MessageResponse{
			Status:  "error",
			Message: "invalid id value",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	pet, err := petSevice.GetPet(id)
	if err != nil {
		response := dto.MessageResponse{
			Status:  "error",
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := dto.PetResponse{
		Status: "success",
		Pet:    dto.PetToPetDTO(pet),
	}
	c.JSON(http.StatusOK, response)
}

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

func UpdatePet(c *gin.Context) {
	petSevice := services.NewPetSevice()
	var petDTO dto.PetDTO
	if err := c.ShouldBindJSON(&petDTO); err != nil {
		response := dto.MessageResponse{
			Status:  "error",
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	pet, err := petSevice.UpdatePet(dto.PetDTOToPet(petDTO))
	if err != nil {
		response := dto.MessageResponse{
			Status:  "error",
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := dto.PetResponse{
		Status: "success",
		Pet:    dto.PetToPetDTO(pet),
	}
	c.JSON(http.StatusOK, response)
}

func DeletePet(c *gin.Context) {
	petSevice := services.NewPetSevice()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response := dto.MessageResponse{
			Status:  "error",
			Message: "invalid id value",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if err = petSevice.DeletePet(id); err != nil {
		response := dto.MessageResponse{
			Status:  "error",
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := dto.MessageResponse{
		Status:  "success",
		Message: "pet deleted",
	}
	c.JSON(http.StatusOK, response)
}

func PetsGeneralStatistics(c *gin.Context) {
	petSevice := services.NewPetSevice()
	statistics, err := petSevice.GetPetsGeneralStatistics()
	if err != nil {
		response := dto.MessageResponse{
			Status:  "error",
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK,
		dto.PetsGeneralStatisticsToPetsGeneralStatisticsDTO(statistics))
}

func GetPetsStatisticsBySpecie(c *gin.Context) {
	petSevice := services.NewPetSevice()
	specie := c.Param("specie")
	statistics, err := petSevice.GetPetsStatisticsBySpecie(specie)
	if err != nil {
		response := dto.MessageResponse{
			Status:  "error",
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK,
		dto.PetsStatisticsBySpecieToPetsStatisticsBySpecieDTO(statistics))
}

func GetPetsBySpecie(c *gin.Context) {
	petSevice := services.NewPetSevice()
	specieParamer := c.Param("specie")
	specieResult, err := petSevice.GetPetsBySpecie(specieParamer)
	if err != nil {
		response := dto.MessageResponse{
			Status:  "error",
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK,
		dto.PetsToPetsDTO(specieResult))
}

func GetPetsByGender(c *gin.Context) {
	petSevice := services.NewPetSevice()
	genderParamer := c.Param("gender")
	petsByGenderResult, err := petSevice.GetPetsByGender(genderParamer)
	if err != nil {
		response := dto.MessageResponse{
			Status:  "error",
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK,
		dto.PetsToPetsDTO(petsByGenderResult))
}

func Help(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"doc/api.html",
		gin.H{},
	)
}