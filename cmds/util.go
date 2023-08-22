package cmds

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pivaldi/db-migration/drivers"
	goose "github.com/pressly/goose/v3"
)

func isStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func gooseRunWithOptions(command string, args []string) {
	var err error

	if !isStringInSlice(dbmConfig.DBConnection.Driver, driversList) {
		var format = "Driver \"%s\" is not supported.\nAvaillable drivers are :\n  %s\n"
		panic(fmt.Sprintf(format, dbmConfig.DBConnection.Driver, strings.Join(driversList, "\n  ")))
	}

	switch dbmConfig.DBConnection.Driver {
	case "sqlite3":
		//  Internally uses the CGo-free port of SQLite: modernc.org/sqlite
		dbmConfig.DBConnection.Driver = "sqlite"
	case "postgres":
		dbmConfig.DBConnection.Driver = "pgx"
	}

	var normalizedDBString = drivers.NormalizeDBString(
		dbmConfig.DBConnection.Driver, dbmConfig.DBConnection.DSN,
		dbmConfig.DBConnection.CertFile, dbmConfig.DBConnection.SSLCert, dbmConfig.DBConnection.SSLKey)

	db, err := goose.OpenDBWithDriver(dbmConfig.DBConnection.Driver, normalizedDBString)
	if err != nil {
		panic(fmt.Sprintf("DSN=%q : %v\n", dbmConfig.DBConnection.DSN, err))
	}

	options := []goose.OptionsFunc{}

	if dbmConfig.DBMigration.AllowMissing {
		options = append(options, goose.WithAllowMissing())
	}

	if dbmConfig.DBMigration.NoVersioning {
		options = append(options, goose.WithNoVersioning())
	}

	if err := goose.RunWithOptions(command, db, dbmConfig.DBMigration.Dir, args, options...); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
