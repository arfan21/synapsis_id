package main

import (
	"os"

	"github.com/arfan21/synapsis_id/cmd/api"
	migration "github.com/arfan21/synapsis_id/cmd/migrate"
	"github.com/urfave/cli/v2"
)

// @title Synapsis ID API
// @version 1.0
// @description This is a sample server cell for Synapsis ID Test API.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.synapsis.id
// @contact.email
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8888
// @BasePath /api
func main() {
	appCli := cli.NewApp()
	appCli.Name = "Synapsis ID Test"
	appCli.Usage = "Synapsis ID Test API"
	appCli.Commands = []*cli.Command{
		migration.Root(),
		api.Serve(),
	}

	if err := appCli.Run(os.Args); err != nil {
		panic(err)
	}
}
