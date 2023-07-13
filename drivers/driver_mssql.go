//go:build !no_mssql
// +build !no_mssql

package drivers

import (
	_ "github.com/denisenkom/go-mssqldb"
)
