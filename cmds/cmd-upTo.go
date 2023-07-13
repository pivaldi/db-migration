package cmds

import (
	"github.com/spf13/cobra"
)

var upToCmd = &cobra.Command{
	Use:              "upTo [VERSION]",
	Short:            "Migrate the DB to a specific VERSION",
	TraverseChildren: true,
	Args:             cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		gooseRunWithOptions("up-to", args)
	},
}
