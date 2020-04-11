package models

import (
	"fmt"
)

type ErrInvalidSliceType struct {
	Model Model
	Value interface{}
}

func (err *ErrInvalidSliceType) Error() string {
	return fmt.Sprintf("Model %s can't use slice of type %T", err.Model.Name(), err.Value)
}

type ErrNotAssociated struct {
	Model Model
	Assoc string
}

func (err *ErrNotAssociated) Error() string {
	return fmt.Sprintf("Model %s has no association named %q", err.Model.Name(), err.Assoc)
}

type ErrInvalidRecord struct {
	Model  Model
	Record Record
}

func (err *ErrInvalidRecord) Error() string {
	return fmt.Sprintf(
		"Invalid record (model %s, type %T) for model %s",
		err.Record.Model().Name(),
		err.Record,
		err.Model.Name(),
	)
}

type ErrInvalidAssociationRecord struct {
	Assoc  Association
	Record Record
}

func (err *ErrInvalidAssociationRecord) Error() string {
	return fmt.Sprintf(
		"Invalid record (model %s, type %T) for association %s",
		err.Record.Model().Name(),
		err.Record,
		AssociationString(err.Assoc),
	)
}
