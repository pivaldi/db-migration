package cmds

import (
	"github.com/spf13/cobra"
)

var downToCmd = &cobra.Command{
	Use:              "downTo [VERSION]",
	Short:            "Roll back the DB migrations to a specific VERSION",
	TraverseChildren: true,
	Args:             cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		gooseRunWithOptions("down-to", args)
	},
}
