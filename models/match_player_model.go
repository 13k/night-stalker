// Code generated by modelgen. DO NOT EDIT.

package models

var _ Record = (*MatchPlayer)(nil)

var MatchPlayerModel Model = (*modelMatchPlayer)(nil)

// uses an anonymous struct to have compile-time checks when using associations
var assocMatchPlayer = struct {
	Match Association
	Hero  Association
}{
	Match: NewAssociation(
		BelongsTo,
		&ModelAssociation{Model: MatchPlayerModel, Name: "Match", PK: "id", FK: "match_id"},
		&ModelAssociation{Model: MatchModel, Name: "MatchPlayer", PK: "id", FK: ""},
	),
	Hero: NewAssociation(
		BelongsTo,
		&ModelAssociation{Model: MatchPlayerModel, Name: "Hero", PK: "id", FK: "hero_id"},
		&ModelAssociation{Model: HeroModel, Name: "MatchPlayer", PK: "id", FK: ""},
	),
}

type modelMatchPlayer struct{}

func (*modelMatchPlayer) Name() string             { return "MatchPlayer" }
func (*modelMatchPlayer) Table() Table             { return MatchPlayerTable }
func (*modelMatchPlayer) NewRecord() Record        { return &MatchPlayer{} }
func (*modelMatchPlayer) NewSlicePtr() interface{} { return &[]*MatchPlayer{} }

func (*modelMatchPlayer) AsRecordSlice(v interface{}) ([]Record, error) {
	sptr, ok := v.(*[]*MatchPlayer)

	if !ok {
		return nil, &ErrInvalidSliceType{Model: MatchPlayerModel, Value: v}
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

func (*modelMatchPlayer) Associations() []Association {
	return []Association{
		assocMatchPlayer.Match,
		assocMatchPlayer.Hero,
	}
}

func (*modelMatchPlayer) Association(name string) (Association, error) {
	switch name {
	case "Match":
		return assocMatchPlayer.Match, nil
	case "Hero":
		return assocMatchPlayer.Hero, nil
	}

	return nil, &ErrNotAssociated{Model: MatchPlayerModel, Assoc: name}
}

func (*MatchPlayer) Model() Model {
	return MatchPlayerModel
}

func (m *MatchPlayer) AssignRecord(record Record) (bool, error) {
	if other, ok := record.(*MatchPlayer); ok {
		return m.Assign(other), nil
	}

	return false, &ErrInvalidRecord{Model: m.Model(), Record: record}
}

func (m *MatchPlayer) AssignPartialRecord(record Record) (bool, error) {
	if other, ok := record.(*MatchPlayer); ok {
		return m.AssignPartial(other), nil
	}

	return false, &ErrInvalidRecord{Model: MatchPlayerModel, Record: record}
}

func (m *MatchPlayer) GetAssocPK(assoc string) (ID, error) {
	return m.ID, nil
}

func (m *MatchPlayer) GetAssocFK(assoc string) (ID, error) {
	switch assoc {
	case "Match":
		return m.MatchID, nil
	case "Hero":
		return m.HeroID, nil
	}

	return 0, &ErrNotAssociated{Model: MatchPlayerModel, Assoc: assoc}
}

func (m *MatchPlayer) SetAssociated(assoc string, records ...Record) error {
	if len(records) == 0 {
		return nil
	}

	switch assoc {
	case "Match":
		r := records[0]

		if mr, ok := r.(*Match); ok {
			m.Match = mr
			return nil
		}

		return &ErrInvalidAssociationRecord{Assoc: assocMatchPlayer.Match, Record: r}
	case "Hero":
		r := records[0]

		if mr, ok := r.(*Hero); ok {
			m.Hero = mr
			return nil
		}

		return &ErrInvalidAssociationRecord{Assoc: assocMatchPlayer.Hero, Record: r}
	}

	return &ErrNotAssociated{Model: MatchPlayerModel, Assoc: assoc}
}