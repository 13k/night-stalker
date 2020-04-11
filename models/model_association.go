package models

import (
	"fmt"

	"golang.org/x/xerrors"
)

type AssociationCardinality uint

const (
	InvalidCardinality AssociationCardinality = iota
	OneToOne
	OneToMany
	ManyToMany
)

func (c AssociationCardinality) String() string {
	switch c {
	case OneToOne:
		return "one-to-one"
	case OneToMany:
		return "one-to-many"
	case ManyToMany:
		return "many-to-many"
	default:
		return "<invalid>"
	}
}

type AssociationType uint

const (
	// BelongsTo represents an one-to-one association where the current model holds the foreign key of
	// the associated model.
	BelongsTo AssociationType = iota
	// HasOne represents an one-to-one association where the associated model holds the foreign key of
	// the current model. It's the counterpart of `BelongsTo`.
	HasOne
	// HasMany represents an one-to-many association where the associated model holds the foreign key
	// of the current model.
	HasMany
)

func (t AssociationType) String() string {
	switch t {
	case BelongsTo:
		return "belongs_to"
	case HasOne:
		return "has_one"
	case HasMany:
		return "has_many"
	default:
		return "<invalid>"
	}
}

func (t AssociationType) Cardinality() AssociationCardinality {
	switch t {
	case BelongsTo:
		return OneToOne
	case HasOne:
		return OneToOne
	case HasMany:
		return OneToMany
	default:
		return InvalidCardinality
	}
}

// ModelAssociation describes one side of an Association.
type ModelAssociation struct {
	Model Model  // Model
	Name  string // Name of the association in Model
	PK    string // Primary key column of the association in Model
	FK    string // Foreign key column of the association in Model
}

// Association describes an unidirectional relation between two models.
type Association interface {
	// Type returns the association type.
	Type() AssociationType
	// Source returns the source model association.
	Source() *ModelAssociation
	// Dest returns the destination model association.
	Dest() *ModelAssociation
	// PK returns the association primary key column, depending on association type.
	//
	// BelongsTo: column used to hold the primary key on the destination model (usually `dest.id`).
	// HasOne:    column used to hold the primary key on the source model (usually `source.id`).
	// HasMany:   column used to hold the primary key on the source model (usually `source.id`).
	PK() Column
	// FK returns the association foreign key column, depending on association type.
	//
	// BelongsTo: column used to hold the foreign key on the source model (usually `source.<association>_id`).
	// HasOne:    column used to hold the foreign key on the destination model (usually `dest.<association>_id`).
	// HasMany:   column used to hold the foreign key on the destination model (usually `dest.<association>_id`).
	FK() Column
	// Col returns the association "join" column, depending on association type.
	//
	// BelongsTo: return value of `PK`.
	// HasOne:    return value of `FK`.
	// HasMany:   return value of `FK`.
	Col() Column
	// CollectIDs collects association IDs from a batch of source records, depending on association
	// type.
	//
	// BelongsTo: ID value corresponding to `FK` (usually the ID stored in `source.<association>_id`).
	// HasOne:    ID value corresponding to `PK` (usually the ID stored in `source.id`).
	// HasMany:   ID value corresponding to `PK` (usually the ID stored in `source.id`).
	//
	// All source records must have the same Model as `Source` (compared by name).
	CollectIDs(sources []Record) ([]ID, error)
	// SetRecords associates a batch of source records with their corresponding destination records.
	//
	// All source records must have the same Model as `Source` and all destination records must have
	// the same Model as `Dest` (compared by name).
	SetRecords(sources []Record, dests []Record) error
}

func AssociationString(a Association) string {
	return fmt.Sprintf(
		"%s.%s %s %s.%s",
		a.Source().Model.Name(),
		a.Source().Name,
		a.Type(),
		a.Dest().Model.Name(),
		a.Dest().Name,
	)
}

var _ Association = (*association)(nil)

type association struct {
	typ  AssociationType
	src  *ModelAssociation
	dest *ModelAssociation
}

func NewAssociation(typ AssociationType, src *ModelAssociation, dest *ModelAssociation) Association {
	return &association{typ: typ, src: src, dest: dest}
}

func NewBelongsToAssociation(src *ModelAssociation, dest *ModelAssociation) Association {
	return NewAssociation(BelongsTo, src, dest)
}

func NewHasOneAssociation(src *ModelAssociation, dest *ModelAssociation) Association {
	return NewAssociation(HasOne, src, dest)
}

func NewHasManyAssociation(src *ModelAssociation, dest *ModelAssociation) Association {
	return NewAssociation(HasMany, src, dest)
}

func (a *association) String() string            { return AssociationString(a) }
func (a *association) Type() AssociationType     { return a.typ }
func (a *association) Source() *ModelAssociation { return a.src }
func (a *association) Dest() *ModelAssociation   { return a.dest }

// BelongsTo: column used to hold the primary key on the destination model (usually `dest.id`).
// HasOne:    column used to hold the primary key on the source model (usually `source.id`).
// HasMany:   column used to hold the primary key on the source model (usually `source.id`).
func (a *association) PK() Column {
	switch a.typ {
	case BelongsTo:
		return a.dest.Model.Table().Col(a.dest.PK)
	case HasOne, HasMany:
		return a.src.Model.Table().Col(a.src.PK)
	default:
		return nil
	}
}

// BelongsTo: column used to hold the foreign key on the source model (usually `source.<association>_id`).
// HasOne:    column used to hold the foreign key on the destination model (usually `dest.<association>_id`).
// HasMany:   column used to hold the foreign key on the destination model (usually `dest.<association>_id`).
func (a *association) FK() Column {
	switch a.typ {
	case BelongsTo:
		return a.src.Model.Table().Col(a.src.FK)
	case HasOne, HasMany:
		return a.dest.Model.Table().Col(a.dest.FK)
	default:
		return nil
	}
}

// BelongsTo: return value of `PK`.
// HasOne:    return value of `FK`.
// HasMany:   return value of `FK`.
func (a *association) Col() Column {
	switch a.typ {
	case BelongsTo:
		return a.PK()
	case HasOne, HasMany:
		return a.FK()
	default:
		return nil
	}
}

func (a *association) getSourceID(s Record) (ID, error) {
	if s.Model().Name() != a.src.Model.Name() {
		return 0, &ErrInvalidAssociationRecord{Assoc: a, Record: s}
	}

	var id ID
	var err error

	switch a.typ {
	case BelongsTo:
		id, err = s.GetAssocFK(a.src.Name)
	case HasOne, HasMany:
		id, err = s.GetAssocPK(a.src.Name)
	}

	if err != nil {
		return 0, xerrors.Errorf("error collecting source ID: %w", err)
	}

	return id, nil
}

func (a *association) getDestID(d Record) (ID, error) {
	if d.Model().Name() != a.dest.Model.Name() {
		return 0, &ErrInvalidAssociationRecord{Assoc: a, Record: d}
	}

	var id ID
	var err error

	switch a.typ {
	case BelongsTo:
		id, err = d.GetAssocPK(a.dest.Name)
	case HasOne, HasMany:
		id, err = d.GetAssocFK(a.dest.Name)
	}

	if err != nil {
		return 0, xerrors.Errorf("error collecting destination ID: %w", err)
	}

	return id, nil
}

// BelongsTo: ID value corresponding to `FK` (usually the ID stored in `source.<association>_id`).
// HasOne:    ID value corresponding to `PK` (usually the ID stored in `source.id`).
// HasMany:   ID value corresponding to `PK` (usually the ID stored in `source.id`).
func (a *association) CollectIDs(sources []Record) ([]ID, error) {
	ids := make([]ID, len(sources))

	for i, s := range sources {
		id, err := a.getSourceID(s)

		if err != nil {
			return nil, xerrors.Errorf("CollectIDs: error collecting sources IDs: %w", err)
		}

		ids[i] = id
	}

	return NewUniqueIDs(ids...), nil
}

func (a *association) SetRecords(sources []Record, dests []Record) error {
	destByID := make(map[ID][]Record, len(dests))

	// group dests by association ID
	for _, d := range dests {
		id, err := a.getDestID(d)

		if err != nil {
			return xerrors.Errorf("SetRecords: error collecting destination IDs: %w", err)
		}

		destByID[id] = append(destByID[id], d)
	}

	// associate sources to dests by association ID
	for _, s := range sources {
		id, err := a.getSourceID(s)

		if err != nil {
			return xerrors.Errorf("SetRecords: error collecting source IDs: %w", err)
		}

		associated := destByID[id]

		if err := s.SetAssociated(a.src.Name, associated...); err != nil {
			return xerrors.Errorf("SetRecords: error associating model records: %w", err)
		}
	}

	return nil
}
