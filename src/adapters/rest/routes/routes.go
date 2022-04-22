package routes

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/oswaldom-code/skaypet-api/src/adapters/rest/controllers"
	"github.com/spf13/viper"
)

//SetupRouter sets up routes
func SetupRouter() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()
	// TODO: add auth middleware
	storeCookie := cookie.NewStore([]byte(viper.GetString("auth.secret")))
	router.Use(sessions.Sessions("session", storeCookie))
	public := router.Group("/v1.0")
	{
		public.GET("/ping", controllers.Pong)
		public.POST("/mascota/crear", controllers.CreatePet)
		public.GET("/mascota/listar", controllers.GetPets)
		public.GET("/mascota/:id", controllers.GetPet)
		public.PATCH("/mascota/:id", controllers.UpdatePet)
		public.DELETE("/mascota/:id", controllers.DeletePet)
	}
	return router
}
