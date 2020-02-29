package routes

import (
	"fmt"

	"github.com/spf13/cobra"

	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	nsweb "github.com/13k/night-stalker/web"
)

var Cmd = &cobra.Command{
	Use:   "routes",
	Short: "List routes",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	log, err := nscmdlog.New()

	if err != nil {
		panic(err)
	}

	defer log.Close()

	app, err := nsweb.New(nsweb.AppOptions{
		Log: log,
	})

	if err != nil {
		log.WithError(err).Fatal("error initializing application")
	}

	for _, route := range app.Routes() {
		fmt.Printf("%7s %s\n", route.Method, route.Path)
	}
}
