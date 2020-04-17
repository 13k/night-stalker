package collections

import (
	nsm "github.com/13k/night-stalker/models"
)

type Teams []*nsm.Team

func (s Teams) KeyByID() map[nsm.ID]*nsm.Team {
	if s == nil {
		return nil
	}

	m := make(map[nsm.ID]*nsm.Team)

	for _, t := range s {
		m[t.ID] = t
	}

	return m
}
