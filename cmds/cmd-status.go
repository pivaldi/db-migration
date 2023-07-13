package cmds

import (
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:              "status",
	Short:            "Dump the migration status for the current DB",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		gooseRunWithOptions("status", args)
	},
}
