// Code generated by modelgen. DO NOT EDIT.

package models

var _ Record = (*LiveMatchStatsBuilding)(nil)

var LiveMatchStatsBuildingModel Model = (*modelLiveMatchStatsBuilding)(nil)

// uses an anonymous struct to have compile-time checks when using associations
var assocLiveMatchStatsBuilding = struct {
	LiveMatchStats Association
}{
	LiveMatchStats: NewAssociation(
		BelongsTo,
		&ModelAssociation{Model: LiveMatchStatsBuildingModel, Name: "LiveMatchStats", PK: "id", FK: "live_match_stats_id"},
		&ModelAssociation{Model: LiveMatchStatsModel, Name: "LiveMatchStatsBuilding", PK: "id", FK: ""},
	),
}

type modelLiveMatchStatsBuilding struct{}

func (*modelLiveMatchStatsBuilding) Name() string             { return "LiveMatchStatsBuilding" }
func (*modelLiveMatchStatsBuilding) Table() Table             { return LiveMatchStatsBuildingTable }
func (*modelLiveMatchStatsBuilding) NewRecord() Record        { return &LiveMatchStatsBuilding{} }
func (*modelLiveMatchStatsBuilding) NewSlicePtr() interface{} { return &[]*LiveMatchStatsBuilding{} }

func (*modelLiveMatchStatsBuilding) AsRecordSlice(v interface{}) ([]Record, error) {
	sptr, ok := v.(*[]*LiveMatchStatsBuilding)

	if !ok {
		return nil, &ErrInvalidSliceType{Model: LiveMatchStatsBuildingModel, Value: v}
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

func (*modelLiveMatchStatsBuilding) Associations() []Association {
	return []Association{
		assocLiveMatchStatsBuilding.LiveMatchStats,
	}
}

func (*modelLiveMatchStatsBuilding) Association(name string) (Association, error) {
	switch name {
	case "LiveMatchStats":
		return assocLiveMatchStatsBuilding.LiveMatchStats, nil
	}

	return nil, &ErrNotAssociated{Model: LiveMatchStatsBuildingModel, Assoc: name}
}

func (*LiveMatchStatsBuilding) Model() Model {
	return LiveMatchStatsBuildingModel
}

func (m *LiveMatchStatsBuilding) AssignRecord(record Record) (bool, error) {
	if other, ok := record.(*LiveMatchStatsBuilding); ok {
		return m.Assign(other), nil
	}

	return false, &ErrInvalidRecord{Model: m.Model(), Record: record}
}

func (m *LiveMatchStatsBuilding) AssignPartialRecord(record Record) (bool, error) {
	if other, ok := record.(*LiveMatchStatsBuilding); ok {
		return m.AssignPartial(other), nil
	}

	return false, &ErrInvalidRecord{Model: LiveMatchStatsBuildingModel, Record: record}
}

func (m *LiveMatchStatsBuilding) GetAssocPK(assoc string) (ID, error) {
	return m.ID, nil
}

func (m *LiveMatchStatsBuilding) GetAssocFK(assoc string) (ID, error) {
	switch assoc {
	case "LiveMatchStats":
		return m.LiveMatchStatsID, nil
	}

	return 0, &ErrNotAssociated{Model: LiveMatchStatsBuildingModel, Assoc: assoc}
}

func (m *LiveMatchStatsBuilding) SetAssociated(assoc string, records ...Record) error {
	if len(records) == 0 {
		return nil
	}

	switch assoc {
	case "LiveMatchStats":
		r := records[0]

		if mr, ok := r.(*LiveMatchStats); ok {
			m.LiveMatchStats = mr
			return nil
		}

		return &ErrInvalidAssociationRecord{Assoc: assocLiveMatchStatsBuilding.LiveMatchStats, Record: r}
	}

	return &ErrNotAssociated{Model: LiveMatchStatsBuildingModel, Assoc: assoc}
}