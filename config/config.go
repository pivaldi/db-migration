package config

// import (
// 	"github.com/spf13/viper"
// )

type DBConnectionT struct {
	DSN      string `mapstructure:"dsn"`
	Driver   string `mapstructure:"driver"`
	CertFile string `mapstructure:"cert-file"`
	SSLCert  string `mapstructure:"ssl-cert"`
	SSLKey   string `mapstructure:"ssl-key"`
}

type DBMigrationT struct {
	Dir          string `mapstructure:"dir"`
	Table        string `mapstructure:"table"`
	Verbose      bool   `mapstructure:"verbose"`
	Sequential   bool   `mapstructure:"sequential"`
	AllowMissing bool   `mapstructure:"allow-missing"`
	NoVersioning bool   `mapstructure:"no-versioning"`
}

type ConfigT struct {
	DBConnection DBConnectionT
	DBMigration  DBMigrationT
}
