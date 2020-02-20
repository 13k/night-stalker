package collections

import (
	"sort"

	"github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protocol"
)

type TVGames []*protocol.CSourceTVGameSmall

// Clean cleans up Source TV games.
//
// * Removes nil games
// * Removes games with invalid MatchId (zero)
// * De-duplicates games with same MatchId
//
// De-duplication: elects the game with highest LastUpdateTime as the valid entry. Elected entries
// are placed in the same position as the first occurrence of duplicated entries.
//
// It returns the slice sorted by descending SortScore.
func (s TVGames) Clean() TVGames {
	if s == nil {
		return nil
	}

	groupByMatchID := make(map[nspb.MatchID]TVGames)

	for _, game := range s {
		if game == nil {
			continue
		}

		if game.GetMatchId() == 0 {
			continue
		}

		groupByMatchID[game.GetMatchId()] = append(groupByMatchID[game.GetMatchId()], game)
	}

	visited := make(map[nspb.MatchID]bool)
	result := make(TVGames, 0, len(s))

	for _, game := range s {
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
