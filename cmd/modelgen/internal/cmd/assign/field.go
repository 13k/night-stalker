package cmdassign

import (
	"fmt"

	g "github.com/13k/night-stalker/cmd/modelgen/internal/parser"
)

type field struct {
	Name string

	op *typeOp
}

func parseFields(s *g.Struct) ([]*field, error) {
	var fs []*field

	for _, f := range s.DBFieldSet() {
		if !f.Exported() {
			continue
		}

		pf, err := newField(f)

		if err != nil {
			return nil, fmt.Errorf("could not parse field %s.%s: %w", s.Name, f.Name(), err)
		}

		fs = append(fs, pf)
	}

	return fs, nil
}

func newField(sf *g.StructField) (*field, error) {
	op, err := parseTypeOp(sf.Type())

	if err != nil {
		return nil, fmt.Errorf("could not parse field type: %w", err)
	}

	f := &field{
		Name: sf.Name(),
		op:   op,
	}

	return f, nil
}

func (f *field) NotZero(v string) string {
	v = fmt.Sprintf("%s.%s", v, f.Name)

	return f.op.NotZero(v)
}

func (f *field) NotEqual(left, right string) string {
	left = fmt.Sprintf("%s.%s", left, f.Name)
	right = fmt.Sprintf("%s.%s", right, f.Name)

	return f.op.NotEqual(left, right)
}

func (f *field) Assign(left, right string) string {
	left = fmt.Sprintf("%s.%s", left, f.Name)
	right = fmt.Sprintf("%s.%s", right, f.Name)

	return f.op.Assign(left, right)
}
