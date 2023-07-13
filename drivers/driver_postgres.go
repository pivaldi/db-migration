//go:build !no_postgres
// +build !no_postgres

package drivers

import (
	_ "github.com/jackc/pgx/v5/stdlib"
)
