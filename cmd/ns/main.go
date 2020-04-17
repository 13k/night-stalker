package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmdroot"
	nscmdmeta "github.com/13k/night-stalker/cmd/ns/internal/meta"

	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func main() {
	meta := &nscmdmeta.Meta{
		Version:  version,
		Revision: revision,
	}

	if err := cmdroot.Execute(meta); err != nil {
		if !errors.Is(err, cmdroot.ErrCommandFailureLogged) {
			fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
		}

		os.Exit(1)
	}
}
