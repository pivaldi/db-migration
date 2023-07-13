package cmds

import (
	"github.com/pivaldi/db-migration/config"
)

const (
	noVersioningDescr = "Apply migration commands with no versioning, in file order, from directory pointed to."
	certfileDescr     = "File path to root CA's certificates in pem format (only support on mysql)."
	directory         = "db-migrations"
)

var (
	dbmConfig = config.ConfigT{}
	cfgFile   = Root.PersistentFlags().String("config", "", "Configuration file to use")
)

func SetConfig(config config.ConfigT) {
	dbmConfig = config
}

func init() {
	Root.PersistentFlags().StringVarP(
		&dbmConfig.DBMigration.Dir, "dir", "d", directory, "Directory with migration files.")

	Root.PersistentFlags().StringVarP(
		&dbmConfig.DBMigration.Table, "table", "", "db_migration", "The DB migration table name.")

	Root.PersistentFlags().BoolVarP(
		&dbmConfig.DBMigration.Verbose, "verbose", "v", false, "Enable verbose mode")

	Root.PersistentFlags().StringVarP(
		&dbmConfig.DBConnection.CertFile, "certfile", "", "", certfileDescr)

	Root.PersistentFlags().BoolVarP(
		&dbmConfig.DBMigration.Sequential, "sequential", "s", false, "Use sequential numbering for new migrations.")

	Root.PersistentFlags().BoolVarP(
		&dbmConfig.DBMigration.AllowMissing, "allow-missing", "a", false, "Applies missing (out-of-order) migrations.")

	Root.PersistentFlags().BoolVarP(
		&dbmConfig.DBMigration.NoVersioning, "no-versioning", "", false, noVersioningDescr)

	Root.PersistentFlags().StringVar(
		&dbmConfig.DBConnection.Driver, "driver", "", "SQL driver to use (see the `example` subcommand for the list).")

	Root.PersistentFlags().StringVar(
		&dbmConfig.DBConnection.DSN, "dsn", "", "The Data Source Name (see the `example` subcommand)")

	Root.MarkFlagsRequiredTogether("driver", "dsn")

	Root.PersistentFlags().StringVar(
		&dbmConfig.DBConnection.SSLCert, "ssl-cert", "", "File path to SSL certificates in pem format (only support on mysql).")

	Root.PersistentFlags().StringVarP(
		&dbmConfig.DBConnection.SSLKey, "ssl-key", "k", "", "File path to SSL key in pem format (only support on mysql).")
}
