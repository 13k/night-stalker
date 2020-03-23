package main

import (
	"fmt"
	"os"

	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmdroot"
	nscmdmeta "github.com/13k/night-stalker/cmd/ns/internal/meta"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	meta := &nscmdmeta.Meta{
		Version:  version,
		Revision: revision,
	}

	if err := cmdroot.Execute(meta); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
		os.Exit(1)
	}
}
