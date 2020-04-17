// Code generated by modelgen. DO NOT EDIT.

package models

var _ Record = (*League)(nil)

var LeagueModel Model = (*modelLeague)(nil)

type modelLeague struct{}

func (*modelLeague) Name() string             { return "League" }
func (*modelLeague) Table() Table             { return LeagueTable }
func (*modelLeague) NewRecord() Record        { return &League{} }
func (*modelLeague) NewSlicePtr() interface{} { return &[]*League{} }

func (*modelLeague) AsRecordSlice(v interface{}) ([]Record, error) {
	sptr, ok := v.(*[]*League)

	if !ok {
		return nil, &ErrInvalidSliceType{Model: LeagueModel, Value: v}
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

func (*modelLeague) Associations() []Association {
	return nil
}

func (*modelLeague) Association(name string) (Association, error) {
	return nil, &ErrNotAssociated{Model: LeagueModel, Assoc: name}
}

func (*League) Model() Model {
	return LeagueModel
}

func (m *League) AssignRecord(record Record) (bool, error) {
	if other, ok := record.(*League); ok {
		return m.Assign(other), nil
	}

	return false, &ErrInvalidRecord{Model: m.Model(), Record: record}
}

func (m *League) AssignPartialRecord(record Record) (bool, error) {
	if other, ok := record.(*League); ok {
		return m.AssignPartial(other), nil
	}

	return false, &ErrInvalidRecord{Model: LeagueModel, Record: record}
}

func (m *League) GetAssocPK(assoc string) (ID, error) {
	return m.ID, nil
}

func (m *League) GetAssocFK(assoc string) (ID, error) {
	return 0, &ErrNotAssociated{Model: LeagueModel, Assoc: assoc}
}

func (m *League) SetAssociated(assoc string, records ...Record) error {
	return &ErrNotAssociated{Model: LeagueModel, Assoc: assoc}
}