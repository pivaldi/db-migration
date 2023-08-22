package cmds

import (
	goose "github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var driversList = []string{
	"postgres", "mysql", "sqlite3", "mssql", "redshift", "tidb", "clickhouse", "vertica",
}

var Root = &cobra.Command{
	Use:              "db-migration [OPTIONS] COMMAND",
	Short:            "Database migration command line",
	TraverseChildren: true,
}

func init() {
	cobra.OnInitialize(initViperConfig)

	if dbmConfig.DBMigration.Verbose {
		goose.SetVerbose(true)
	}

	if dbmConfig.DBMigration.Sequential {
		goose.SetSequential(true)
	}

	goose.SetTableName(dbmConfig.DBMigration.Table)

	Root.AddCommand(versionCmd)
	Root.AddCommand(exampleCmd)
	Root.AddCommand(initCmd)
	Root.AddCommand(createCmd)
	Root.AddCommand(fixCmd)
	Root.AddCommand(upCmd)
	Root.AddCommand(upByOneCmd)
	Root.AddCommand(statusCmd)
	Root.AddCommand(upToCmd)
	Root.AddCommand(donwCmd)
	Root.AddCommand(downToCmd)
}

// Execute executes the root command.
func Execute() error {
	return Root.Execute()
}
