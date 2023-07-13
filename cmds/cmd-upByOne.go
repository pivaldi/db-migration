package cmds

import "github.com/spf13/cobra"

var upByOneCmd = &cobra.Command{
	Use:              "upByOne",
	Short:            "Migrate the DB up by 1",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		gooseRunWithOptions("up-by-one", args)
	},
}
