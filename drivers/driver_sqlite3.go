//go:build !no_sqlite3 && !(windows && arm64)
// +build !no_sqlite3
// +build !windows !arm64

package drivers

import (
	_ "modernc.org/sqlite"
)
