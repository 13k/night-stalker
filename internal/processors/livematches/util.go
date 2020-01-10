package livematches

import (
	"sort"

	"github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protocol"
)

// cleanResponseGames cleans response games.
//
// * Removes nil games
// * Removes games with invalid MatchId (zero)
// * De-duplicates games with same MatchId
//
// De-duplication: elects the game with highest LastUpdateTime as the valid entry. Elected entries
// are placed in the same position as the first occurrence of duplicated entries.
//
// It returns the slice sorted by descending SortScore.
func cleanResponseGames(games []*protocol.CSourceTVGameSmall) []*protocol.CSourceTVGameSmall {
	if games == nil {
		return nil
	}

	groupByMatchID := make(map[nspb.MatchID][]*protocol.CSourceTVGameSmall)

	for _, game := range games {
		if game == nil {
			continue
		}

		if game.GetMatchId() == 0 {
			continue
		}

		groupByMatchID[game.GetMatchId()] = append(groupByMatchID[game.GetMatchId()], game)
	}

	visited := make(map[nspb.MatchID]bool)
	result := make([]*protocol.CSourceTVGameSmall, 0, len(games))

	for _, game := range games {
		if game == nil {
			continue
		}

		if game.GetMatchId() == 0 {
			continue
		}

		if visited[game.GetMatchId()] {
			continue
		}

		group := groupByMatchID[game.GetMatchId()]

		if len(group) > 1 {
			for _, g := range group {
				if g.GetLastUpdateTime() > game.GetLastUpdateTime() {
					game = g
				}
			}
		}

		result = append(result, game)
		visited[game.GetMatchId()] = true
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].GetSortScore() > result[j].GetSortScore()
	})

	return result
}
