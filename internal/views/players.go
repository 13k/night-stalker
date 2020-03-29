package views

import (
	"sort"

	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

func NewPlayers(data PlayersData) ([]*nspb.Player, error) {
	if len(data) == 0 {
		return nil, nil
	}

	players := make([]*nspb.Player, 0, len(data))

	for _, playerData := range data {
		player, err := NewPlayer(playerData)

		if err != nil {
			err = xerrors.Errorf("error creating Player view: %w", err)
			return nil, err
		}

		players = append(players, player)
	}

	return players, nil
}

func NewSortedPlayers(data PlayersData) ([]*nspb.Player, error) {
	players, err := NewPlayers(data)

	if err != nil {
		err = xerrors.Errorf("error creating Player views: %w", err)
		return nil, err
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].AccountId < players[j].AccountId
	})

	return players, nil
}
