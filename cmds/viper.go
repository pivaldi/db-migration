package cmds

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var runtimeViper = viper.New()

func initViperConfig() {
	viperBindPFlags()
	initViperFileConfig()
	runtimeViper.AutomaticEnv()
}

func initViperFileConfig() {
	if *cfgFile != "" {
		if _, err := os.Stat(*cfgFile); err != nil {
			panic(err)
		}

		runtimeViper.SetConfigFile(*cfgFile)

		if err := runtimeViper.ReadInConfig(); err != nil {
			panic(err)
		}

		log.Println("Viper uses config file : ", runtimeViper.ConfigFileUsed())

		if err := runtimeViper.Unmarshal(&dbmConfig); err != nil {
			panic(err)
		}
	}
}

func viperBindPFlag(flag string) {
	if err := runtimeViper.BindPFlag(flag, Root.PersistentFlags().Lookup(flag)); err != nil {
		panic(err)
	}
}

func viperBindPFlags() {
	viperBindPFlag("dir")
	viperBindPFlag("table")
	viperBindPFlag("verbose")
	viperBindPFlag("certfile")
	viperBindPFlag("sequential")
	viperBindPFlag("allow-missing")
	viperBindPFlag("ssl-key")
	viperBindPFlag("no-versioning")
	viperBindPFlag("driver")
	viperBindPFlag("dsn")
}
