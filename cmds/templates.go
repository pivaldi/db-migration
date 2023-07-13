package cmds

import "text/template"

var sqlMigrationTemplate = template.Must(template.New("goose.sql-migration").Parse(`
-- Documentation can be found here: https://pressly.github.io/goose
-- The -- +goose NO TRANSACTION directive may be added to the top of the file to run statements
-- outside a transaction.

-- +goose Up
-- +goose StatementBegin

SELECT 'up SQL query';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

SELECT 'down SQL query';

-- +goose StatementEnd
`))

var goMigrationTemplate = template.Must(template.New("goose.go-migration").Parse(`package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(up{{.CamelName}}, down{{.CamelName}})
}

func up{{.CamelName}}(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func down{{.CamelName}}(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
`))
