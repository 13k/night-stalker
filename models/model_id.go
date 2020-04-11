package models

// ID is the type of a model's ID field.
//
// It seems stupid to have getter and setter methods, but it's useful for embedding in model structs
// and the structs "inherit" the methods.
type ID uint64

func (id ID) GetID() ID {
	return id
}

func (id *ID) SetID(i ID) {
	*id = i
}

// IDs is a slice of ID
type IDs []ID

func NewUniqueIDs(ids ...ID) IDs {
	if len(ids) == 0 {
		return nil
	}

	s := make([]ID, 0, len(ids))
	index := make(map[ID]struct{}, len(ids))

	for _, id := range ids {
		if _, ok := index[id]; !ok {
			index[id] = struct{}{}
			s = append(s, id)
		}
	}

	return s
}
