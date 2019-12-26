package main

import (
	"fmt"
	"os"

	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmdroot"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	if err := cmdroot.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}
