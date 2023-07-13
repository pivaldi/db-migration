package main

import (
	"log"

	"github.com/pivaldi/db-migration/cmds"
)

func main() {
	if err := cmds.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
