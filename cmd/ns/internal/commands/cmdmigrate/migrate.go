package cmdmigrate

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/markbates/pkger"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdhttpfsd "github.com/13k/night-stalker/cmd/ns/internal/httpfsd"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	v "github.com/13k/night-stalker/cmd/ns/internal/viper"
	nsdb "github.com/13k/night-stalker/internal/db"
	nslog "github.com/13k/night-stalker/internal/logger"
)

var Cmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	RunE:  run,
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

func run(cmd *cobra.Command, args []string) error {
	log := nscmdlog.Instance()

	defer log.Close()

	if flagGenerate != "" {
		return runGenerate(log)
	}

	db, err := nscmddb.Connect(log)

	if err != nil {
		return xerrors.Errorf("error connecting to database: %w", err)
	}

	defer db.Close()

	if flagPrint {
		return runPrint(db, log)
	}

	return runMigrate(db, log)
}

func runGenerate(log *nslog.Logger) error {
	if flagPath == "" {
		return xerrors.New("missing migrations path")
	}

	version := flagVersion

	if version == 0 {
		version = uint(time.Now().Unix())
	}

	if err := os.MkdirAll(flagPath, os.ModePerm); err != nil {
		return xerrors.Errorf("error creating migrations directory: %w", err)
	}

	versionGlob := fmt.Sprintf("%d_*", version)
	versionGlob = filepath.Join(flagPath, versionGlob)

	if matches, err := filepath.Glob(versionGlob); err != nil {
		return xerrors.Errorf("error generating migration: %w", err)
	} else if len(matches) > 0 {
		return xerrors.Errorf("duplicate migration version %d", version)
	}

	for _, direction := range []string{"up", "down"} {
		basename := fmt.Sprintf("%d_%s.%s.sql", version, flagGenerate, direction)
		filename := filepath.Join(flagPath, basename)
		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

		if err != nil {
			return xerrors.Errorf("error generating migration: %w", err)
		}

		defer f.Close()

		log.WithOFields(
			"version", version,
			"direction", direction,
			"file", filename,
		).Info("created")
	}

	return nil
}

func newMigrate(db *nsdb.DB, log *nslog.Logger) (*migrate.Migrate, error) {
	pkgerRoot, err := pkger.Open("/migrations")

	if err != nil {
		return nil, xerrors.Errorf("error opening migrations directory: %w", err)
	}

	srcDriver, err := nscmdhttpfsd.New(pkgerRoot, "/")

	if err != nil {
		return nil, xerrors.Errorf("error creating httpfs: %w", err)
	}

	dbDriver, err := postgres.WithInstance(db.SQLDB(), &postgres.Config{})

	if err != nil {
		return nil, xerrors.Errorf("error creating postgres driver: %w", err)
	}

	m, err := migrate.NewWithInstance("pkger", srcDriver, v.GetString(v.KeyDbDriver), dbDriver)

	if err != nil {
		return nil, xerrors.Errorf("error creating migrator: %w", err)
	}

	m.Log = log.MigrateLogger()

	return m, nil
}

func runPrint(db *nsdb.DB, log *nslog.Logger) error {
	m, err := newMigrate(db, log)

	if err != nil {
		return xerrors.Errorf("error creating migrator: %w", err)
	}

	defer m.Close()

	version, dirty, err := m.Version()

	if err != nil {
		if xerrors.Is(err, migrate.ErrNilVersion) {
			log.WithField("version", "nil").Info("no migration applied yet")
			return nil
		}

		return xerrors.Errorf("error fetching version: %w", err)
	}

	l := log.WithField("version", version)

	if dirty {
		l.Warn("database needs manual fix!")
	} else {
		l.Info("fetched migration version")
	}

	return nil
}

func runMigrate(db *nsdb.DB, log *nslog.Logger) error {
	if flagVersion != 0 && flagSteps != 0 {
		return xerrors.New("version and steps options are mutually exclusive")
	}

	if flagForceVersion && flagVersion == 0 {
		return xerrors.New("force option must be used with version")
	}

	m, err := newMigrate(db, log)

	if err != nil {
		return xerrors.Errorf("error creating migrator: %w", err)
	}

	defer m.Close()

	var l *nslog.Logger = log
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

	if err := fn(); err != nil {
		if xerrors.Is(err, migrate.ErrNoChange) {
			log.Info("already up-to-date")
			return nil
		}

		return xerrors.Errorf("error migrating database: %w", err)
	}

	log.Info("done")

	return nil
}
