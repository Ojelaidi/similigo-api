package similigo_api

import (
	"github.com/Ojelaidi/similigo-api/internal/api"
	"github.com/spf13/cobra"
)

func Start() *cobra.Command {
	cmds := &cobra.Command{
		Use: "api",
		Run: run,
	}

	_ = cmds.PersistentFlags()
	cmds.AddCommand()

	return cmds
}

func run(cmd *cobra.Command, args []string) {
	api.New().Run()
}
