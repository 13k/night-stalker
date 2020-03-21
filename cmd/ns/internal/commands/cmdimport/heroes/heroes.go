package heroes

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cobra"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	nskv "github.com/13k/night-stalker/internal/kv"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
	"github.com/13k/night-stalker/models"
)

var (
	kvNameAliasesSepRE = regexp.MustCompile(`[,;]`)
	kvRoleSepRE        = regexp.MustCompile(`[,]`)
	kvRoleLevelsSepRE  = regexp.MustCompile(`[,]`)
)

var Cmd = &cobra.Command{
	Use:   "heroes <npc_heroes.txt>",
	Short: "Import heroes from KeyValues file",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	log, err := nscmdlog.New()

	if err != nil {
		panic(err)
	}

	defer log.Close()

	if len(args) < 1 {
		log.Fatal("KeyValues file argument is required")
	}

	log.Info("parsing KeyValues file ...")

	kvRoot, err := nskv.ParseFile(args[0])

	if err != nil {
		log.WithError(err).Fatal("error parsing KeyValues file")
	}

	if kvRoot.Key() != "DOTAHeroes" {
		log.WithField("root", kvRoot.Key()).Fatal("invalid KeyValues file: wrong root node key")
	}

	kvHeroes, err := kvRoot.Children()

	if err != nil {
		log.WithError(err).Fatal("invalid KeyValues file")
	}

	var heroes []*models.Hero

	for _, node := range kvHeroes {
		l := log.WithField("key", node.Key())

		var enabled bool

		enabled, err = kvIsEnabledHeroNode(node)

		if err != nil {
			l.WithError(err).Fatal("error parsing hero KeyValues")
		}

		if !enabled {
			continue
		}

		var hero *models.Hero

		hero, err = kvParseHeroNode(node)

		if err != nil {
			l.WithError(err).Fatal("error parsing hero KeyValues")
		}

		if hero != nil {
			heroes = append(heroes, hero)
		}
	}

	db, err := nscmddb.Connect()

	if err != nil {
		log.WithError(err).Fatal("error connecting to database")
	}

	defer db.Close()

	log.Info("importing heroes ...")

	for _, hero := range heroes {
		l := log.WithField("name", hero.Name)
		result := db.Where(&models.Hero{ID: hero.ID}).Assign(hero).FirstOrCreate(hero)

		if err := result.Error; err != nil {
			l.WithError(err).Fatal("error creating or updating hero")
		}

		l.Info("imported")
	}

	log.Info("done")
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
		return false, err
	}

	return enabled == 1, nil
}

func kvParseHeroNode(node *nskv.Node) (*models.Hero, error) {
	if !kvIsHeroNode(node) {
		return nil, fmt.Errorf("node '%s' is not a hero node", node.Key())
	}

	hero := &models.Hero{Name: node.Key()}

	id, err := node.ChildAsInt64("HeroID", false)

	if err != nil {
		return nil, err
	}

	hero.ID = nspb.HeroID(id)

	hero.LocalizedName, err = node.ChildAsString("workshop_guide_name", false, true)

	if err != nil {
		return nil, err
	}

	aliases, err := node.ChildAsStringArray("NameAliases", kvNameAliasesSepRE, true, true, true)

	if err != nil {
		return nil, err
	}

	if len(aliases) > 0 {
		hero.Aliases = aliases
	}

	roles, err := node.ChildAsStringArray("Role", kvRoleSepRE, true, true, false)

	if err != nil {
		return nil, err
	}

	if len(roles) > 0 {
		rolesPb := make(nssql.HeroRoles, len(roles))

		for i, roleStr := range roles {
			roleEnumKey := "HERO_ROLE_" + strings.ToUpper(roleStr)
			roleInt, ok := nspb.HeroRole_value[roleEnumKey]

			if !ok {
				return nil, fmt.Errorf("invalid Role item value '%s' in node '%s'", roleStr, node.Key())
			}

			rolesPb[i] = nspb.HeroRole(roleInt)
		}

		hero.Roles = rolesPb
	}

	roleLevels, err := node.ChildAsInt64Array("Rolelevels", kvRoleLevelsSepRE, true)

	if err != nil {
		return nil, err
	}

	if len(roleLevels) > 0 {
		hero.RoleLevels = roleLevels
	}

	hero.Complexity, err = node.ChildAsInt("Complexity", true)

	if err != nil {
		return nil, err
	}

	hero.Legs, err = node.ChildAsInt("Legs", true)

	if err != nil {
		return nil, err
	}

	attackCapStr, err := node.ChildAsString("AttackCapabilities", false, true)

	if err != nil {
		return nil, err
	}

	attackCapInt, ok := nspb.DotaUnitCap_value[attackCapStr]

	if !ok {
		return nil, fmt.Errorf("invalid AttackCapabilities property value '%s' in node '%s'", attackCapStr, node.Key())
	}

	hero.AttackCapabilities = nspb.DotaUnitCap(attackCapInt)

	attrPrimaryStr, err := node.ChildAsString("AttributePrimary", false, true)

	if err != nil {
		return nil, err
	}

	attrPrimaryInt, ok := nspb.DotaAttribute_value[attrPrimaryStr]

	if !ok {
		return nil, fmt.Errorf("invalid AttributePrimary property value '%s' in node '%s'", attrPrimaryStr, node.Key())
	}

	hero.AttributePrimary = nspb.DotaAttribute(attrPrimaryInt)

	return hero, nil
}
