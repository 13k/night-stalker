package views

import (
	"sort"

	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

func NewMatches(data MatchesData) ([]*nspb.Match, error) {
	if len(data) == 0 {
		return nil, nil
	}

	matches := make([]*nspb.Match, 0, len(data))

	for _, matchData := range data {
		match, err := NewMatch(matchData)

		if err != nil {
			err = xerrors.Errorf("error creating Match view: %w", err)
			return nil, err
		}

		matches = append(matches, match)
	}

	return matches, nil
}

func NewSortedMatches(data MatchesData) ([]*nspb.Match, error) {
	matches, err := NewMatches(data)

	if err != nil {
		err = xerrors.Errorf("error creating Match views: %w", err)
		return nil, err
	}

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].MatchId < matches[j].MatchId
	})

	return matches, nil
}
