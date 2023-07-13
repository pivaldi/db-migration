package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var examplesDescr = `
Examples:
    db-migration --config=./test/config.yml create
    db-migration --driver=sqlite3 --dsn=./foo.db status
    db-migration --driver=sqlite3 --dsn=./foo.db create --name=add_some_column --type=sql
    db-migration --driver=sqlite3 --dsn=./foo.db up

    db-migration --driver=postgres --dsn="user=postgres password=postgres dbname=postgres sslmode=disable" status
    db-migration --driver=mysql --dsn="user:password@/dbname?parseTime=true" status
    db-migration --driver=redshift --dsn="postgres://user:password@qwerty.us-east-1.redshift.amazonaws.com:5439/db" status
    db-migration --driver=tidb --dsn="user:password@/dbname?parseTime=true" status
    db-migration --driver=mssql --dsn="sqlserver://user:password@dbname:1433?database=master" status
    db-migration --driver=clickhouse --dsn="tcp://127.0.0.1:9000" status
    db-migration --driver=vertica --dsn="vertica://user:password@localhost:5433/dbname?connection_load_balance=1" status

Drivers:
%s`

var exampleCmd = &cobra.Command{
	Use:   "example",
	Short: "Show examples of commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(examplesDescr, driversList)
	},
}
