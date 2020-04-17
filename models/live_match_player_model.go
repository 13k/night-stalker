// Code generated by modelgen. DO NOT EDIT.

package models

var _ Record = (*LiveMatchPlayer)(nil)

var LiveMatchPlayerModel Model = (*modelLiveMatchPlayer)(nil)

// uses an anonymous struct to have compile-time checks when using associations
var assocLiveMatchPlayer = struct {
	LiveMatch Association
	Match     Association
	Hero      Association
}{
	LiveMatch: NewAssociation(
		BelongsTo,
		&ModelAssociation{Model: LiveMatchPlayerModel, Name: "LiveMatch", PK: "id", FK: "live_match_id"},
		&ModelAssociation{Model: LiveMatchModel, Name: "LiveMatchPlayer", PK: "id", FK: ""},
	),
	Match: NewAssociation(
		BelongsTo,
		&ModelAssociation{Model: LiveMatchPlayerModel, Name: "Match", PK: "id", FK: "match_id"},
		&ModelAssociation{Model: MatchModel, Name: "LiveMatchPlayer", PK: "id", FK: ""},
	),
	Hero: NewAssociation(
		BelongsTo,
		&ModelAssociation{Model: LiveMatchPlayerModel, Name: "Hero", PK: "id", FK: "hero_id"},
		&ModelAssociation{Model: HeroModel, Name: "LiveMatchPlayer", PK: "id", FK: ""},
	),
}

type modelLiveMatchPlayer struct{}

func (*modelLiveMatchPlayer) Name() string             { return "LiveMatchPlayer" }
func (*modelLiveMatchPlayer) Table() Table             { return LiveMatchPlayerTable }
func (*modelLiveMatchPlayer) NewRecord() Record        { return &LiveMatchPlayer{} }
func (*modelLiveMatchPlayer) NewSlicePtr() interface{} { return &[]*LiveMatchPlayer{} }

func (*modelLiveMatchPlayer) AsRecordSlice(v interface{}) ([]Record, error) {
	sptr, ok := v.(*[]*LiveMatchPlayer)

	if !ok {
		return nil, &ErrInvalidSliceType{Model: LiveMatchPlayerModel, Value: v}
	}

	if sptr == nil || *sptr == nil {
		return nil, nil
	}

	rs := make([]Record, len(*sptr))

	for i, r := range *sptr {
		rs[i] = r
	}

	return rs, nil
}

func (*modelLiveMatchPlayer) Associations() []Association {
	return []Association{
		assocLiveMatchPlayer.LiveMatch,
		assocLiveMatchPlayer.Match,
		assocLiveMatchPlayer.Hero,
	}
}

func (*modelLiveMatchPlayer) Association(name string) (Association, error) {
	switch name {
	case "LiveMatch":
		return assocLiveMatchPlayer.LiveMatch, nil
	case "Match":
		return assocLiveMatchPlayer.Match, nil
	case "Hero":
		return assocLiveMatchPlayer.Hero, nil
	}

	return nil, &ErrNotAssociated{Model: LiveMatchPlayerModel, Assoc: name}
}

func (*LiveMatchPlayer) Model() Model {
	return LiveMatchPlayerModel
}

func (m *LiveMatchPlayer) AssignRecord(record Record) (bool, error) {
	if other, ok := record.(*LiveMatchPlayer); ok {
		return m.Assign(other), nil
	}

	return false, &ErrInvalidRecord{Model: m.Model(), Record: record}
}

func (m *LiveMatchPlayer) AssignPartialRecord(record Record) (bool, error) {
	if other, ok := record.(*LiveMatchPlayer); ok {
		return m.AssignPartial(other), nil
	}

	return false, &ErrInvalidRecord{Model: LiveMatchPlayerModel, Record: record}
}

func (m *LiveMatchPlayer) GetAssocPK(assoc string) (ID, error) {
	return m.ID, nil
}

func (m *LiveMatchPlayer) GetAssocFK(assoc string) (ID, error) {
	switch assoc {
	case "LiveMatch":
		return m.LiveMatchID, nil
	case "Match":
		return m.MatchID, nil
	case "Hero":
		return m.HeroID, nil
	}

	return 0, &ErrNotAssociated{Model: LiveMatchPlayerModel, Assoc: assoc}
}

func (m *LiveMatchPlayer) SetAssociated(assoc string, records ...Record) error {
	if len(records) == 0 {
		return nil
	}

	switch assoc {
	case "LiveMatch":
		r := records[0]

		if mr, ok := r.(*LiveMatch); ok {
			m.LiveMatch = mr
			return nil
		}

		return &ErrInvalidAssociationRecord{Assoc: assocLiveMatchPlayer.LiveMatch, Record: r}
	case "Match":
		r := records[0]

		if mr, ok := r.(*Match); ok {
			m.Match = mr
			return nil
		}

		return &ErrInvalidAssociationRecord{Assoc: assocLiveMatchPlayer.Match, Record: r}
	case "Hero":
		r := records[0]

		if mr, ok := r.(*Hero); ok {
			m.Hero = mr
			return nil
		}

		return &ErrInvalidAssociationRecord{Assoc: assocLiveMatchPlayer.Hero, Record: r}
	}

	return &ErrNotAssociated{Model: LiveMatchPlayerModel, Assoc: assoc}
}