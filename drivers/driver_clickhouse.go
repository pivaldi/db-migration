//go:build !no_clickhouse
// +build !no_clickhouse

package drivers

import (
	_ "github.com/ClickHouse/clickhouse-go/v2"
)
