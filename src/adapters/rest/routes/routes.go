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
	storeCookie := cookie.NewStore([]byte(viper.GetString("auth.secret")))
	router.Use(sessions.Sessions("session", storeCookie))
	public := router.Group("/v1.0")
	{
		public.GET("/ping", controllers.GetPets)
	}

	//v1 := router.Group("/v1.0")
	////v1.Use(ValidateSessionMiddleware)
	//{
	//	v1.GET("/blog/post/:id", controllers.BlogPost)
	//
	//}

	return router
}
