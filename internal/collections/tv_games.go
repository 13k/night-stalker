package collections

import (
	"sort"

	"github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

type TVGames []*protocol.CSourceTVGameSmall

func (s TVGames) MatchIDs() MatchIDs {
	if s == nil {
		return nil
	}

	matchIDs := make(MatchIDs, len(s))

	for i, game := range s {
		matchIDs[i] = nspb.MatchID(game.GetMatchId())
	}

	return matchIDs
}

func (s TVGames) FindIndexByMatchID(matchID nspb.MatchID) int {
	for i, game := range s {
		if nspb.MatchID(game.GetMatchId()) == matchID {
			return i
		}
	}

	return -1
}

func (s TVGames) GroupByMatchID() map[nspb.MatchID]TVGames {
	if s == nil {
		return nil
	}

	m := make(map[nspb.MatchID]TVGames)

	for _, game := range s {
		if game == nil {
			continue
		}

		if game.GetMatchId() == 0 {
			continue
		}

		matchID := nspb.MatchID(game.GetMatchId())
		m[matchID] = append(m[matchID], game)
	}

	return m
}

func (s TVGames) Shift() (TVGames, *protocol.CSourceTVGameSmall) {
	if len(s) == 0 {
		return s, nil
	}

	return s[1:], s[0]
}

func (s TVGames) Pop() (TVGames, *protocol.CSourceTVGameSmall) {
	if len(s) == 0 {
		return s, nil
	}

	return s[:len(s)-1], s[len(s)-1]
}

func (s TVGames) Remove(i int) (TVGames, *protocol.CSourceTVGameSmall) {
	if i < 0 || i >= len(s) {
		return s, nil
	}

	if i == 0 {
		return s.Shift()
	}

	if i == len(s)-1 {
		return s.Pop()
	}

	game := s[i]

	copy(s[i:], s[i+1:])
	s[len(s)-1] = nil
	s = s[:len(s)-1]

	return s, game
}

func (s TVGames) RemoveByMatchID(matchID nspb.MatchID) (TVGames, *protocol.CSourceTVGameSmall) {
	return s.Remove(s.FindIndexByMatchID(matchID))
}

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

	byMatchID := s.GroupByMatchID()
	visited := make(map[nspb.MatchID]bool)
	result := make(TVGames, 0, len(s))

	for _, game := range s {
		if game == nil {
			continue
		}

		if game.GetMatchId() == 0 {
			continue
		}

		matchID := nspb.MatchID(game.GetMatchId())

		if visited[matchID] {
			continue
		}

		group := byMatchID[matchID]

		if len(group) > 1 {
			for _, g := range group {
				if g.GetLastUpdateTime() > game.GetLastUpdateTime() {
					game = g
				}
			}
		}

		result = append(result, game)
		visited[matchID] = true
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].GetSortScore() > result[j].GetSortScore()
	})

	return result
}
