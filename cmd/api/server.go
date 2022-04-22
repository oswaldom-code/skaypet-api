package cmd

import (
	"log"

	"github.com/oswaldom-code/skaypet-api/pkg/config"
	persistence "github.com/oswaldom-code/skaypet-api/src/adapters/persistence/repository"
	"github.com/oswaldom-code/skaypet-api/src/adapters/rest/routes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(serveCmdNew)
}

var serveCmdNew = &cobra.Command{
	Use:   "server",
	Short: "Spin up the webserver that hosts the API",
	Long:  `The web server hosts the API, and manages the authentication middleware`,
	Run: func(cmd *cobra.Command, args []string) {
		NewServer()
	},
}

func NewServer() {
	// Load persistence config
	persistence.New(config.GetDBConfig())

	//setup routes
	r := routes.SetupRouter()
	err := r.Run(":" + viper.GetString("server.port"))
	if err != nil {
		log.Println("Error: ", err.Error())
	}
}
