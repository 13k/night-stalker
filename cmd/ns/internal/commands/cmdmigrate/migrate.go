package cmdmigrate

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jinzhu/gorm"
	"github.com/markbates/pkger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdhttpfsd "github.com/13k/night-stalker/cmd/ns/internal/httpfsd"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	v "github.com/13k/night-stalker/cmd/ns/internal/viper"
	nslog "github.com/13k/night-stalker/internal/logger"
)

var Cmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	Run:   run,
}

var (
	flagPath         string
	flagGenerate     string
	flagPrint        bool
	flagVersion      uint
	flagForceVersion bool
	flagSteps        int
)

func init() {
	Cmd.Flags().StringVarP(&flagPath, "path", "p", "migrations", "migrations path")
	Cmd.Flags().StringVarP(&flagGenerate, "gen", "g", "", "generate migration")
	Cmd.Flags().BoolVarP(&flagPrint, "print", "P", false, "print the currently active migration version")
	Cmd.Flags().UintVarP(&flagVersion, "version", "V", 0, "migration version")
	Cmd.Flags().BoolVarP(&flagForceVersion, "force", "f", false, "force migration version (must be used with --version)")
	Cmd.Flags().IntVarP(&flagSteps, "steps", "s", 0, "migrate N steps (can be negative)")
}

func run(cmd *cobra.Command, args []string) {
	log, err := nscmdlog.New()

	if err != nil {
		panic(err)
	}

	defer log.Close()

	if flagGenerate != "" {
		runGenerate(log)
		return
	}

	db, err := nscmddb.Connect()

	if err != nil {
		log.WithError(err).Fatal("error connecting to database")
	}

	defer db.Close()

	if flagPrint {
		runPrint(db, log)
		return
	}

	runMigrate(db, log)
}

func runGenerate(log *nslog.Logger) {
	if flagPath == "" {
		log.Fatal("missing migrations path")
	}

	version := flagVersion

	if version == 0 {
		version = uint(time.Now().Unix())
	}

	if err := os.MkdirAll(flagPath, os.ModePerm); err != nil {
		log.WithError(err).Fatal("error creating migrations directory")
	}

	versionGlob := fmt.Sprintf("%d_*", version)
	versionGlob = filepath.Join(flagPath, versionGlob)

	if matches, err := filepath.Glob(versionGlob); err != nil {
		log.WithError(err).Fatal("error generating migration")
	} else if len(matches) > 0 {
		log.WithField("version", version).Fatal("duplicate migration version")
	}

	for _, direction := range []string{"up", "down"} {
		basename := fmt.Sprintf("%d_%s.%s.sql", version, flagGenerate, direction)
		filename := filepath.Join(flagPath, basename)
		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

		if err != nil {
			log.WithError(err).Fatal("error generating migration")
		}

		f.Close()

		log.WithField("file", filename).Info("created")
	}
}

func newMigrate(db *gorm.DB, log *nslog.Logger) *migrate.Migrate {
	pkgerRoot, err := pkger.Open("/migrations")

	if err != nil {
		log.WithError(err).Fatal("error opening pkger root")
	}

	srcDriver, err := nscmdhttpfsd.New(pkgerRoot, "/")

	if err != nil {
		log.WithError(err).Fatal("error creating httpfs source driver")
	}

	dbDriver, err := postgres.WithInstance(db.DB(), &postgres.Config{})

	if err != nil {
		log.WithError(err).Fatal("error creating postgres db driver")
	}

	m, err := migrate.NewWithInstance("pkger", srcDriver, v.GetString(v.KeyDbDriver), dbDriver)

	if err != nil {
		log.WithError(err).Fatal("error creating migrator")
	}

	m.Log = log.MigrateLogger()

	return m
}

func runPrint(db *gorm.DB, log *nslog.Logger) {
	m := newMigrate(db, log)

	defer m.Close()

	version, dirty, err := m.Version()

	switch err {
	case nil:
		l := log.WithField("version", version)

		if dirty {
			l.Warn("database needs manual fix!")
		} else {
			l.Info("fetched migration version")
		}
	case migrate.ErrNilVersion:
		log.WithField("version", "nil").Info("no migration applied yet")
	default:
		log.WithError(err).Error("error fetching version")
	}
}

func runMigrate(db *gorm.DB, log *nslog.Logger) {
	if flagVersion != 0 && flagSteps != 0 {
		log.Fatal("version and steps options are mutually exclusive")
	}

	if flagForceVersion && flagVersion == 0 {
		log.Fatal("force option must be used with version")
	}

	m := newMigrate(db, log)

	defer m.Close()

	var l logrus.FieldLogger = log
	var fn func() error

	switch {
	case flagVersion != 0:
		l = log.WithField("version", flagVersion)

		if flagForceVersion {
			l = l.WithField("force", true)
			fn = func() error { return m.Force(int(flagVersion)) }
		} else {
			fn = func() error { return m.Migrate(flagVersion) }
		}
	case flagSteps != 0:
		l = log.WithField("steps", flagSteps)
		fn = func() error { return m.Steps(flagSteps) }
	default:
		fn = m.Up
	}

	l.Info("migrating database")

	switch err := fn(); err {
	case nil:
		log.Info("done")
	case migrate.ErrNoChange:
		log.Info("already up-to-date")
	default:
		log.WithError(err).Fatal("error migrating database")
	}
}
