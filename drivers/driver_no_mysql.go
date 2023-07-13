//go:build no_mysql
// +build no_mysql

package drivers

func NormalizeDBString(driver string, dsn string, certfile string, sslcert string, sslkey string) string {
	return dsn
}
