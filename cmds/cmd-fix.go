package cmds

import (
	"log"

	goose "github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	Use:              "fix",
	Short:            "Apply sequential ordering to migrations.",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		if err := goose.Run("fix", nil, dbmConfig.DBMigration.Dir, args...); err != nil {
			log.Fatalf("%v", err)
		}
	},
}
