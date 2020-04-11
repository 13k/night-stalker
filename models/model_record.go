package models

type Record interface {
	// Model returns the record's Model.
	Model() Model
	// GetID returns the record's ID.
	GetID() ID
	// SetID sets the record's ID.
	SetID(id ID)
	// AssignRecord assigns all fields in the given Record (via type assertion) and returns true, with
	// no error, if any field has changed.
	//
	// If the given record is of invalid type, it returns false and an error of type *ErrInvalidRecord.
	AssignRecord(record Record) (bool, error)
	// AssignPartialRecord assigns all non-zero valued fields in the given Record (via type assertion)
	// and returns true, with no error, if any field has changed.
	//
	// If the given record is of invalid type, it returns false and an error of type *ErrInvalidRecord.
	AssignPartialRecord(record Record) (bool, error)
	// GetAssocPK ...
	GetAssocPK(assoc string) (ID, error)
	// GetAssocFK ...
	GetAssocFK(assoc string) (ID, error)
	// SetAssociated ...
	SetAssociated(assoc string, records ...Record) error
}
