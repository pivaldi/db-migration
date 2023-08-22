package cmds

import "github.com/spf13/cobra"

var donwCmd = &cobra.Command{
	Use:              "down",
	Short:            "Roll back the DB migrations by 1.",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		gooseRunWithOptions("down", args)
	},
}
