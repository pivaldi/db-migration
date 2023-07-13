package cmds

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"

	goose "github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

// initDir will create a directory with an empty SQL migration file.
func gooseInit(dir string) error {
	_, err := os.Stat(dir)
	createDir := true

	switch {
	case errors.Is(err, fs.ErrNotExist):
	case err == nil, errors.Is(err, fs.ErrExist):
		// return fmt.Errorf("Directory already exists : %s", dir)
		createDir = false
	default:
		return err
	}

	if createDir {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}

		fmt.Printf("Directory '%s' created…\n", dir)

	}

	return goose.CreateWithTemplate(nil, dir, sqlMigrationTemplate, "initial", "sql")
}

var initCmd = &cobra.Command{
	Use:              "init",
	Short:            "Create the migrations' directory and an empty initial SQL migration file.",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		if err := gooseInit(dbmConfig.DBMigration.Dir); err != nil {
			log.Fatalf("%v", err)
		}
	},
}
