# Cobra Wrapper of Goose Database Migration Tool

This repository provides a [spf13/Cobra](https://github.com/spf13/cobra)
wrapper of the database migration tool
[pressly/Goose](https://github.com/pressly/goose) with the power capabilities'
configuration of [spf13/viper](https://github.com/spf13/viper).


## Notes

Some normalization have been added to the Goose commands :

* Goose syntax is `goose [OPTIONS] DRIVER DBSTRING`  
  `db-Migration` syntax is  
  `db-migration [GLOBAL OPTIONS] COMMAND [COMMAND OPTIONS] [COMMAND ARGS]`.  
  So `DRIVER` and `DBSTRING` must be passed as parameters instead as arguments like this :  
  `db-migration --driver=sqlite3 --dsn=./test.db up`.  
  Note that `DBSTRING` has been renamed `dsn`…
* The original Goose `create` command syntax is  
  `goose [OPTIONS] DRIVER DBSTRING create NAME [go|sql]`.  
  With `db-Migration` the syntax is :  
  * `db-migration create` that use the default name `migration` and
    default `sql` format ;
  * `db-migration create --name=MY_NAME --type=go` to use the custom
    name `MY_NAME` and `go` format ;
* The default migrations directory is `db-migrations` and can be
  overwrote with the option `--dir` or `-d`
* All options' cli can be overwrote by a configuration file provided
  by the option `--config=_PATH_/config.XXX` thanks to the use of
  the [spf13/viper](https://github.com/spf13/viper) module.

## Usage in an existing Cobra CLI

A simplest usage in an existing Cobra CLI can be

```go
package main

import (
	"github.com/spf13/cobra"

	migrationCmds "github.com/pivaldi/db-migration/cmds"
)

var Cmd = &cobra.Command{
	Use:              "mycli [OPTIONS] COMMAND",
	Short:            "mycli description",
	TraverseChildren: true,
}

func init() {
	Cmd.AddCommand(migrationCmds.Root)
}

func main() {
	_ = Cmd.Execute()
}

```

To rename the `db-migration` command name to `migration` :


```go
package main

import (
	"github.com/spf13/cobra"

	migrationCmds "github.com/pivaldi/db-migration/cmds"
)

var Cmd = &cobra.Command{
	Use:              "mycli [OPTIONS] COMMAND",
	Short:            "mycli Description",
	TraverseChildren: true,
}

var migationCmd = &cobra.Command{
	Use:              "migration [OPTIONS] COMMAND",
	Short:            "Database Migration Manager",
	TraverseChildren: true,
}


func init() {
	Cmd.AddCommand(migationCmds.Root.Commands()...)
}

func main() {
	_ = Cmd.Execute()
}

```

To inject a database configuration of you own app, simply use the
`cmds.SetConfig` as this :


```go
package cmd

import (
	"github.com/pivaldi/db-migration/cmds"
	"ovya.fr/csite/app/config"
)

var Root = cmds.Root

func init() {
	dbConfig := config.Get().Database
	cmds.SetConfig(dbConfig)
}
```

## Generated documentation

```text
❯ ./db-migration
Database migration command line

Usage:
  db-migration [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  create      Creates new migration file with the current timestamp.
  example     Show examples of commands
  fix         Apply sequential ordering to migrations.
  help        Help about any command
  init        Create the migrations' directory and an empty initial SQL migration file.
  status      Dump the migration status for the current DB
  up          Migrate the DB to the most recent version available.
  upByOne     Migrate the DB up by 1
  upTo        Migrate the DB to a specific VERSION
  version     Print the database migration version

Flags:
  -a, --allow-missing     Applies missing (out-of-order) migrations.
      --certfile string   File path to root CA's certificates in pem format (only support on mysql).
      --config string     Configuration file to use
  -d, --dir string        Directory with migration files. (default "db-migrations")
      --driver example    SQL driver to use (see the example subcommand for the list).
      --dsn example       The Data Source Name (see the example subcommand)
  -h, --help              help for db-migration
      --no-versioning     Apply migration commands with no versioning, in file order, from directory pointed to.
  -s, --sequential        Use sequential numbering for new migrations.
      --ssl-cert string   File path to SSL certificates in pem format (only support on mysql).
  -k, --ssl-key string    File path to SSL key in pem format (only support on mysql).
      --table string      The DB migration table name. (default "db_migration")
  -v, --verbose           Enable verbose mode

Use "db-migration [command] --help" for more information about a command.
```
