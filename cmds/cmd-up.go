package cmds

import "github.com/spf13/cobra"

var upCmd = &cobra.Command{
	Use:              "up",
	Short:            "Migrate the DB to the most recent version available.",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		gooseRunWithOptions("up", args)
	},
}
