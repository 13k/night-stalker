package heroes

import (
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	nscol "github.com/13k/night-stalker/internal/collections"
	nskv "github.com/13k/night-stalker/internal/kv"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
	nsm "github.com/13k/night-stalker/models"
)

var (
	kvNameAliasesSepRE = regexp.MustCompile(`[,;]`)
	kvRoleSepRE        = regexp.MustCompile(`[,]`)
	kvRoleLevelsSepRE  = regexp.MustCompile(`[,]`)
)

var Cmd = &cobra.Command{
	Use:   "heroes <npc_heroes.txt>",
	Short: "Import heroes from KeyValues file",
	RunE:  run,
}

func run(cmd *cobra.Command, args []string) error {
	log := nscmdlog.Instance()

	defer log.Close()

	if len(args) < 1 {
		return cmd.Usage()
	}

	log.Info("parsing KeyValues file ...")

	kvRoot, err := nskv.ParseFile(args[0])

	if err != nil {
		return xerrors.Errorf("error parsing KeyValues file: %w", err)
	}

	if kvRoot.Key() != "DOTAHeroes" {
		return xerrors.Errorf("invalid KeyValues file: wrong root node key %q", kvRoot.Key())
	}

	kvHeroes, err := kvRoot.Children()

	if err != nil {
		return xerrors.Errorf("error parsing KeyValues file: %w", err)
	}

	heroes, err := kvParseHeroesNodes(kvHeroes)

	if err != nil {
		return xerrors.Errorf("error parsing KeyValues file: %w", err)
	}

	db, err := nscmddb.Connect(log)

	if err != nil {
		return xerrors.Errorf("error connecting to database: %w", err)
	}

	defer db.Close()

	log.Info("importing heroes ...")

	for _, hero := range heroes {
		q := db.
			Q().
			Select().
			Eq(nsm.HeroTable.PK(), hero.ID)

		created, err := db.M().Upsert(cmd.Context(), hero, q)

		if err != nil {
			return xerrors.Errorf("error upserting hero: %w", err)
		}

		msg := "updated"

		if created {
			msg = "imported"
		}

		log.WithOFields(
			"id", hero.ID,
			"name", hero.Name,
		).Info(msg)
	}

	log.Info("done")

	return nil
}

func kvIsHeroNode(node *nskv.Node) bool {
	return strings.HasPrefix(node.Key(), "npc_dota_hero_")
}

func kvIsEnabledHeroNode(node *nskv.Node) (bool, error) {
	if !kvIsHeroNode(node) {
		return false, nil
	}

	enabled, err := node.ChildAsInt64("Enabled", true)

	if err != nil {
		return false, xerrors.Errorf("invalid child %q: %w", "Enabled", err)
	}

	return enabled == 1, nil
}

func kvParseHeroesNodes(nodes []*nskv.Node) (nscol.Heroes, error) {
	var heroes nscol.Heroes

	for _, node := range nodes {
		enabled, err := kvIsEnabledHeroNode(node)

		if err != nil {
			return nil, xerrors.Errorf("error parsing node %s: %w", node.Key(), err)
		}

		if !enabled {
			continue
		}

		hero, err := kvParseHeroNode(node)

		if err != nil {
			return nil, xerrors.Errorf("error parsing node %q: %w", node.Key(), err)
		}

		if hero != nil {
			heroes = append(heroes, hero)
		}
	}

	return heroes, nil
}

func kvParseHeroNode(node *nskv.Node) (*nsm.Hero, error) {
	if !kvIsHeroNode(node) {
		return nil, xerrors.New("not a hero node")
	}

	hero := &nsm.Hero{
		Name: node.Key(),
	}

	id, err := node.ChildAsInt64("HeroID", false)

	if err != nil {
		return nil, xerrors.Errorf("invalid child %q: %w", "HeroID", err)
	}

	hero.ID = nsm.ID(id)

	hero.LocalizedName, err = node.ChildAsString("workshop_guide_name", false, true)

	if err != nil {
		return nil, xerrors.Errorf("invalid child %q: %w", "workshop_guide_name", err)
	}

	aliases, err := node.ChildAsStringArray("NameAliases", kvNameAliasesSepRE, true, true, true)

	if err != nil {
		return nil, xerrors.Errorf("invalid child %q: %w", "NameAliases", err)
	}

	if len(aliases) > 0 {
		hero.Aliases = aliases
	}

	roles, err := node.ChildAsStringArray("Role", kvRoleSepRE, true, true, false)

	if err != nil {
		return nil, xerrors.Errorf("invalid child %q: %w", "Role", err)
	}

	if len(roles) > 0 {
		rolesPb := make(nssql.HeroRoles, len(roles))

		for i, roleStr := range roles {
			roleEnumKey := "HERO_ROLE_" + strings.ToUpper(roleStr)
			roleInt, ok := nspb.HeroRole_value[roleEnumKey]

			if !ok {
				return nil, xerrors.Errorf("unknown Role item value %q", roleStr)
			}

			rolesPb[i] = nspb.HeroRole(roleInt)
		}

		hero.Roles = rolesPb
	}

	roleLevels, err := node.ChildAsInt64Array("Rolelevels", kvRoleLevelsSepRE, true)

	if err != nil {
		return nil, xerrors.Errorf("error parsing child %q: %w", "Rolelevels", err)
	}

	if len(roleLevels) > 0 {
		hero.RoleLevels = roleLevels
	}

	hero.Complexity, err = node.ChildAsInt("Complexity", true)

	if err != nil {
		return nil, xerrors.Errorf("invalid child %q: %w", "Complexity", err)
	}

	hero.Legs, err = node.ChildAsInt("Legs", true)

	if err != nil {
		return nil, xerrors.Errorf("invalid child %q: %w", "Legs", err)
	}

	attackCapStr, err := node.ChildAsString("AttackCapabilities", false, true)

	if err != nil {
		return nil, xerrors.Errorf("invalid child %q: %w", "AttackCapabilities", err)
	}

	attackCapInt, ok := nspb.DotaUnitCap_value[attackCapStr]

	if !ok {
		return nil, xerrors.Errorf("invalid child %q: unknown value %q", "AttackCapabilities", attackCapStr)
	}

	hero.AttackCapabilities = nspb.DotaUnitCap(attackCapInt)

	attrPrimaryStr, err := node.ChildAsString("AttributePrimary", false, true)

	if err != nil {
		return nil, xerrors.Errorf("invalid child %q: %w", "AttributePrimary", err)
	}

	attrPrimaryInt, ok := nspb.DotaAttribute_value[attrPrimaryStr]

	if !ok {
		return nil, xerrors.Errorf("invalid child %q: unknown value %q", "AttributePrimary", attrPrimaryStr)
	}

	hero.AttributePrimary = nspb.DotaAttribute(attrPrimaryInt)

	return hero, nil
}
