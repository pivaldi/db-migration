package cmds

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:              "version",
	Short:            "Print the database migration version",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		gooseRunWithOptions("version", args)
	},
}
