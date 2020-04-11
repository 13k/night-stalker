// Code generated by modelgen. DO NOT EDIT.

package models

var _ Record = (*Team)(nil)

var TeamModel Model = (*modelTeam)(nil)

type modelTeam struct{}

func (*modelTeam) Name() string             { return "Team" }
func (*modelTeam) Table() Table             { return TeamTable }
func (*modelTeam) NewRecord() Record        { return &Team{} }
func (*modelTeam) NewSlicePtr() interface{} { return &[]*Team{} }

func (*modelTeam) AsRecordSlice(v interface{}) ([]Record, error) {
	sptr, ok := v.(*[]*Team)

	if !ok {
		return nil, &ErrInvalidSliceType{Model: TeamModel, Value: v}
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

func (*modelTeam) Associations() []Association {
	return nil
}

func (*modelTeam) Association(name string) (Association, error) {
	return nil, &ErrNotAssociated{Model: TeamModel, Assoc: name}
}

func (*Team) Model() Model {
	return TeamModel
}

func (m *Team) AssignRecord(record Record) (bool, error) {
	if other, ok := record.(*Team); ok {
		return m.Assign(other), nil
	}

	return false, &ErrInvalidRecord{Model: m.Model(), Record: record}
}

func (m *Team) AssignPartialRecord(record Record) (bool, error) {
	if other, ok := record.(*Team); ok {
		return m.AssignPartial(other), nil
	}

	return false, &ErrInvalidRecord{Model: TeamModel, Record: record}
}

func (m *Team) GetAssocPK(assoc string) (ID, error) {
	return m.ID, nil
}

func (m *Team) GetAssocFK(assoc string) (ID, error) {
	return 0, &ErrNotAssociated{Model: TeamModel, Assoc: assoc}
}

func (m *Team) SetAssociated(assoc string, records ...Record) error {
	return &ErrNotAssociated{Model: TeamModel, Assoc: assoc}
}
