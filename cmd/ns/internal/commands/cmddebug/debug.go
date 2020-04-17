package cmddebug

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/docker/go-units"
	"github.com/markbates/pkger"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	nscol "github.com/13k/night-stalker/internal/collections"
	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

var Cmd = &cobra.Command{
	Use:   "debug",
	Short: "Debug something",
	RunE:  debug,
}

var CmdPkger = &cobra.Command{
	Use:   "pkger",
	Short: "List all pkger embedded files",
	RunE:  debugPkger,
}

func init() {
	Cmd.AddCommand(CmdPkger)
}

func dumpf(format string, values ...interface{}) { //nolint: unused
	dumps := make([]interface{}, len(values))

	for i, v := range values {
		dumps[i] = spew.Sdump(v)
	}

	fmt.Printf(format, dumps...)
}

func debug(cmd *cobra.Command, args []string) error {
	log := nscmdlog.Instance()

	defer log.Close()

	db, err := nscmddb.Connect(log)

	if err != nil {
		return xerrors.Errorf("error connecting to database: %w", err)
	}

	defer db.Close()

	dbl := nsdbda.NewLoader(db)

	pm, err := dbl.PlayerMatchesData(cmd.Context(), &nsdbda.PlayerMatchesParams{AccountID: 94054712})

	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	log.WithField("matches", len(pm.MatchesData)).Info("player matches")

	hm, err := dbl.HeroMatchesData(cmd.Context(), &nsdbda.HeroMatchesParams{HeroID: 38})

	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	log.WithField("matches", len(hm.MatchesData)).Info("hero matches")

	matchIDs := nscol.MatchIDs{
		5326126028,
		5326085094,
		5326125506,
		5326087691,
	}

	finishedMatchIDs, err := dbl.FindMatchIDs(cmd.Context(), nsdbda.MatchFilters{
		MatchIDs: matchIDs,
	})

	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	finishedByID := make(map[nspb.MatchID]struct{}, len(finishedMatchIDs))

	for _, id := range finishedMatchIDs {
		finishedByID[id] = struct{}{}
	}

	for _, id := range matchIDs {
		_, finished := finishedByID[id]

		log.WithOFields(
			"match_id", id,
			"finished", finished,
		).Info("is match finished?")
	}

	lms, err := dbl.LiveMatchStats(cmd.Context(), nsdbda.LiveMatchStatsFilters{
		MatchIDs: matchIDs,
		Latest:   3,
	})

	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	log.WithField("count", len(lms)).Info("LiveMatchStats")

	err = db.M().Eagerload(cmd.Context(), "LiveMatch", lms.Records()...)

	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	for _, s := range lms {
		var lmID nsm.ID

		if s.LiveMatch != nil {
			lmID = s.LiveMatch.ID
		}

		log.WithOFields(
			".LiveMatchID", s.LiveMatchID,
			".LiveMatch.ID", lmID,
			"eq", s.LiveMatchID == lmID,
		).Info("LiveMatchStats -> LiveMatch")
	}

	err = db.M().Eagerload(cmd.Context(), "Players", lms.Records()...)

	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	for _, s := range lms {
		log.WithOFields(
			"ID", s.ID,
			"count", len(s.Players),
		).Info("LiveMatchStats -> LiveMatchStatsPlayer")
	}

	log.Info("hero tx begin")

	tx, err := db.Begin(cmd.Context(), nil)

	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	hero := &nsm.Hero{
		ID:            666,
		Name:          "fake_hero",
		LocalizedName: "Fake Hero",
	}

	created, err := tx.M().Upsert(
		cmd.Context(),
		hero,
		tx.Q().Select().Eq(nsm.HeroTable.PK(), hero.ID),
	)

	if err != nil {
		log.WithError(err).Error("upsert error")
	}

	log.WithOFields(
		"created", created,
		"id", hero.ID,
		"name", hero.Name,
		"localized_name", hero.LocalizedName,
		"slug", hero.Slug,
		"created_at", hero.CreatedAt,
		"updated_at", hero.UpdatedAt,
	).Info("hero")

	hero = &nsm.Hero{
		ID:   666,
		Legs: 666,
	}

	created, err = tx.M().Upsert(
		cmd.Context(),
		hero,
		tx.Q().Select().Eq(nsm.HeroTable.PK(), hero.ID),
	)

	if err != nil {
		log.WithError(err).Error("upsert error")
	}

	log.WithOFields(
		"created", created,
		"id", hero.ID,
		"name", hero.Name,
		"localized_name", hero.LocalizedName,
		"slug", hero.Slug,
		"created_at", hero.CreatedAt,
		"updated_at", hero.UpdatedAt,
	).Info("hero")

	log.Info("hero tx rollback")

	if err = tx.Rollback(); err != nil {
		log.WithError(err).Error("rollback error")
	}

	log.Info("player tx begin")

	tx, err = db.Begin(cmd.Context(), nil)

	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	player := &nsm.Player{
		AccountID: 666,
		Name:      "Fake Player",
	}

	created, err = tx.M().Upsert(
		cmd.Context(),
		player,
		tx.Q().Select().Eq(nsm.PlayerTable.Col("account_id"), player.AccountID),
	)

	if err != nil {
		log.WithError(err).Error("upsert error")
	}

	log.WithOFields(
		"created", created,
		"id", player.ID,
		"account_id", player.AccountID,
		"name", player.Name,
		"created_at", hero.CreatedAt,
		"updated_at", hero.UpdatedAt,
	).Info("player")

	player = &nsm.Player{
		AccountID: 666,
		Name:      "You got SCAMED!",
	}

	created, err = tx.M().Upsert(
		cmd.Context(),
		player,
		tx.Q().Select().Eq(nsm.PlayerTable.Col("account_id"), player.AccountID),
	)

	if err != nil {
		log.WithError(err).Error("upsert error")
	}

	log.WithOFields(
		"created", created,
		"id", player.ID,
		"account_id", player.AccountID,
		"name", player.Name,
		"created_at", hero.CreatedAt,
		"updated_at", hero.UpdatedAt,
	).Info("player")

	log.Info("player tx rollback")

	if err := tx.Rollback(); err != nil {
		log.WithError(err).Error("rollback error")
	}

	return nil
}

func debugPkger(cmd *cobra.Command, args []string) error {
	log := nscmdlog.Instance()

	defer log.Close()

	err := pkger.Walk("/", func(p string, i os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		log.WithOFields(
			"is_dir", i.IsDir(),
			"size", i.Size(),
			"size_h", units.BytesSize(float64(i.Size())),
		).Info(p)

		return nil
	})

	if err != nil {
		return xerrors.Errorf("error walking pkger tree: %w", err)
	}

	return nil
}
