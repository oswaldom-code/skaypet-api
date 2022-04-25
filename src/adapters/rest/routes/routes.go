package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oswaldom-code/skaypet-api/pkg/config"
	"github.com/oswaldom-code/skaypet-api/src/adapters/rest/controllers"
	"github.com/spf13/viper"
)

//SetupRouter sets up routes
func SetupRouter() *gin.Engine {
	gin.SetMode(viper.GetString("server.mode"))
	router := gin.Default()
	router.Static("/home_admin", "./html/static")
	templatesPath:=config.GetProjectPath()+"/src/adapters/rest/html/templates/**/*"
	router.LoadHTMLGlob(templatesPath)

	api := router.Group("/api/v1.0")
	{
		api.GET("/ping", controllers.Pong)
		api.GET("/help", controllers.Help)
		api.POST("/mascota/crear", controllers.CreatePet)
		api.GET("/mascotas", controllers.GetPets)
		api.GET("/mascota/:id", controllers.GetPet)
		api.PUT("/mascota/actualizar", controllers.UpdatePet)
		api.DELETE("/mascota/:id", controllers.DeletePet)
		api.GET("/mascotas/especie/:specie", controllers.GetPetsBySpecie)
		api.GET("/mascotas/genero/:gender", controllers.GetPetsByGender)
		api.GET("/estadisticas/mascotas", controllers.PetsGeneralStatistics)
		api.GET("/estadisticas/mascotas/especie/:specie", controllers.GetPetsStatisticsBySpecie)
	}
	return router
}
