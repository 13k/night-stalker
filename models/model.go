package models

type Model interface {
	// Name returns the model's name.
	Name() string
	// Table returns the model's Table.
	Table() Table
	// NewRecord returns a new Record.
	NewRecord() Record
	// NewSlicePtr returns a pointer to a slice of concrete model records (&[]T{}).
	NewSlicePtr() interface{}
	// AsRecordSlice converts a value returned by NewSlicePtr to a slice of Record.
	//
	// If the given value has invalid type, it returns an error of type *ErrInvalidSliceType.
	AsRecordSlice(interface{}) ([]Record, error)
	// Associations returns the model's associations.
	Associations() []Association
	// Association returns the model's association with the given name.
	//
	// If the model has no association with the given name, it returns an error of type
	// *ErrNotAssociated.
	Association(name string) (Association, error)
}
