package commands

import (
	"github.com/fajarihsan21/go-backend/src/configs/serve"
	database "github.com/fajarihsan21/go-backend/src/database/gorm"
	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Use:   "Golang Backend",
	Short: "Simple API with Golang",
}

func init() {
	initCommand.AddCommand(serve.InitServe)
	initCommand.AddCommand(database.MigrateCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}
