package cmdapi

import (
	"fmt"

	mg "github.com/13k/night-stalker/cmd/modelgen/internal/parser"
)

type association struct {
	Type string

	SrcName string

	srcPKField string
	srcFKField string

	destName    string
	destPKField string
	destFKField string

	srcStruct  *mg.Struct
	destStruct *mg.Struct
}

func (a *association) IsOneToOne() bool {
	return a.Type == "BelongsTo" || a.Type == "HasOne"
}

func (a *association) IsOneToMany() bool {
	return a.Type == "HasMany"
}

func (a *association) SrcModel() string {
	return a.srcStruct.Name
}

func (a *association) DestModel() string {
	return a.destStruct.Name
}

// By convention, dest association name is src model
func (a *association) DestName() string {
	if a.destName != "" {
		return a.destName
	}

	return a.SrcModel()
}

// By convention, source PK field name is "ID"
func (a *association) SrcPKField() string {
	if a.srcPKField != "" {
		return a.srcPKField
	}

	return "ID"
}

// By convention, source FK field name is "<AssociationName>ID"
func (a *association) SrcFKField() string {
	if a.srcFKField != "" {
		return a.srcFKField
	}

	return fmt.Sprintf("%sID", a.SrcName)
}

func (a *association) SrcPKCol() string {
	return a.srcStruct.DBFieldColMap()[a.SrcPKField()]
}

func (a *association) SrcFKCol() string {
	return a.srcStruct.DBFieldColMap()[a.SrcFKField()]
}

// By convention, destination PK field name is "ID"
func (a *association) DestPKField() string {
	if a.destPKField != "" {
		return a.destPKField
	}

	return "ID"
}

// By convention, destination FK field name is "<AssociationName>ID"
func (a *association) DestFKField() string {
	if a.destFKField != "" {
		return a.destFKField
	}

	return fmt.Sprintf("%sID", a.DestName())
}

func (a *association) DestPKCol() string {
	return a.destStruct.DBFieldColMap()[a.DestPKField()]
}

func (a *association) DestFKCol() string {
	return a.destStruct.DBFieldColMap()[a.DestFKField()]
}

func parseModelAssociations(s *mg.Struct) ([]*association, error) {
	var assocs []*association

	for _, f := range s.ModelFieldSet() {
		tag := f.ModelStructTag()
		assoc := &association{
			SrcName:   f.Name(),
			srcStruct: s,
		}

		switch tag.Type {
		case "belongs_to":
			assoc.Type = "BelongsTo"
		case "has_one":
			assoc.Type = "HasOne"
		case "has_many":
			assoc.Type = "HasMany"
		default:
			return nil, fmt.Errorf("unknown association type %q for field %s.%s", tag.Type, s.Name, f.Name())
		}

		assoc.destStruct = s.FindSiblingByType(f.Type())

		if assoc.destStruct == nil {
			return nil, fmt.Errorf(
				"could not find associated model struct for field %s.%s (%s)",
				s.Name,
				f.Name(),
				f.Type(),
			)
		}

		for key, val := range tag.Params {
			switch key {
			case "pk":
				// pk:<Field>
				switch assoc.Type {
				case "BelongsTo":
					assoc.destPKField = val
				case "HasOne", "HasMany":
					assoc.srcPKField = val
				}
			case "fk":
				// fk:<Field>
				switch assoc.Type {
				case "BelongsTo":
					assoc.srcFKField = val
				case "HasOne", "HasMany":
					assoc.destFKField = val
				}
			case "source":
				// source:<DestAssocName> (only for has_one and has_many)
				switch assoc.Type {
				case "HasOne", "HasMany":
					assoc.destName = val
				}
			}
		}

		assocs = append(assocs, assoc)
	}

	return assocs, nil
}
