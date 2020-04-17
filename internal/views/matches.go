package views

import (
	"sort"

	"golang.org/x/xerrors"

	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

func NewMatches(data nsdbda.MatchesData) ([]*nspb.Match, error) {
	if len(data) == 0 {
		return nil, nil
	}

	matches := make([]*nspb.Match, 0, len(data))

	for _, matchData := range data {
		match, err := NewMatch(matchData)

		if err != nil {
			return nil, xerrors.Errorf("error creating Match view: %w", err)
		}

		matches = append(matches, match)
	}

	return matches, nil
}

func NewSortedMatches(data nsdbda.MatchesData) ([]*nspb.Match, error) {
	matches, err := NewMatches(data)

	if err != nil {
		return nil, xerrors.Errorf("error creating Match views: %w", err)
	}

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].MatchId < matches[j].MatchId
	})

	return matches, nil
}
