package cmds

import (
	"log"
	"text/template"

	goose "github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var createType = "sql"
var createName = "migration"

var createCmd = &cobra.Command{
	Use:              "create",
	Short:            "Creates new migration file with the current timestamp.",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			template *template.Template
		)

		switch createType {
		case "sql":
			template = sqlMigrationTemplate
		case "go":
			template = goMigrationTemplate
		default:
			log.Fatalf("Unknown migration type : %q", createType)
		}

		err := goose.CreateWithTemplate(nil, dbmConfig.DBMigration.Dir, template, createName, createType)

		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

func init() {
	createCmd.Flags().StringVarP(&createType, "type", "t", createType, "[sql|go] : Create sql or go migration.")
	createCmd.Flags().StringVarP(&createName, "name", "n", createName, "Suffixe file name")

}
