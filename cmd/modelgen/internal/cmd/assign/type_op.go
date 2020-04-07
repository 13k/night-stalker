package cmdassign

import (
	"fmt"
	"go/types"
	"reflect"
)

var (
	boolTypeOp   = &typeOp{kind: reflect.Bool}
	numTypeOp    = &typeOp{kind: reflect.Int}
	strTypeOp    = &typeOp{kind: reflect.String}
	structTypeOp = &typeOp{kind: reflect.Struct}

	bytesTypeOp = &typeOp{
		kind: reflect.Slice,
		pkgs: map[string]string{"bytes": ""},
		neq:  `!bytes.Equal(%s, %s)`,
	}
)

type typeOp struct {
	kind   reflect.Kind
	pkgs   map[string]string
	assign string
	neq    string
	nzero  string
}

func parseTypeOp(t types.Type) (*typeOp, error) {
	var op *typeOp

	switch tt := t.(type) {
	case *types.Basic:
		switch tt.Kind() {
		case types.Float32, types.Float64,
			types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
			types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64:
			op = numTypeOp
		case types.Bool:
			op = boolTypeOp
		case types.String:
			op = strTypeOp
		}
	case *types.Named:
		switch tt.Obj().Pkg().Path() {
		case "database/sql":
			switch tt.Obj().Name() {
			case "NullTime":
				op = &typeOp{
					kind:  reflect.Struct,
					pkgs:  map[string]string{"github.com/13k/night-stalker/internal/sql": "nssql"},
					neq:   `!nssql.NullTimeEqual(%s, %s)`,
					nzero: "!nssql.NullTimeIsZero(%s)",
				}
			}
		case "time":
			switch tt.Obj().Name() {
			case "Duration":
				op = numTypeOp
			case "Time":
				op = structTypeOp
			}
		case "github.com/13k/night-stalker/internal/protobuf/protocol":
			op = numTypeOp
		case "github.com/13k/night-stalker/internal/sql":
			switch tt.Obj().Name() {
			case "Base64Bytes":
				op = bytesTypeOp
			case "AbilityIDs", "AccountIDs", "HeroRoles", "ItemIDs":
				op = &typeOp{
					kind: reflect.Slice,
					pkgs: map[string]string{"github.com/13k/night-stalker/internal/sql": "nssql"},
					neq:  `!nssql.IntArrayEqual(%s, %s)`,
				}
			}
		case "github.com/13k/night-stalker/models":
			switch tt.Obj().Name() {
			case "ID":
				op = numTypeOp
			}
		case "github.com/faceit/go-steam/steamid":
			switch tt.Obj().Name() {
			case "SteamId":
				op = numTypeOp
			}
		case "github.com/lib/pq":
			switch tt.Obj().Name() {
			case "Int64Array":
				op = &typeOp{
					kind: reflect.Slice,
					neq:  "!IntsEqual(%s, %s)",
				}
			case "StringArray":
				op = &typeOp{
					kind: reflect.Slice,
					neq:  "!StringsEqual(%s, %s)",
				}
			}
		}
	case *types.Slice:
		switch ttt := tt.Elem().(type) {
		case *types.Basic:
			switch ttt.Kind() {
			case types.Byte: // []byte
				op = bytesTypeOp
			}
		}
	}

	if op == nil {
		return nil, fmt.Errorf("unknown type %s", t.String())
	}

	if err := op.init(); err != nil {
		return nil, fmt.Errorf("error initializing typeOp: %w", err)
	}

	return op, nil
}

func (op *typeOp) init() error {
	if op.assign == "" {
		switch op.kind {
		case reflect.Bool:
			op.assign = `%s = %s`
		case reflect.Float32, reflect.Float64,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			op.assign = `%s = %s`
		case reflect.String:
			op.assign = `%s = %s`
		case reflect.Struct:
			op.assign = `%s = %s`
		case reflect.Slice:
			op.assign = `%s = %s`
		}
	}

	if op.neq == "" {
		switch op.kind {
		case reflect.Bool:
			op.neq = `%s != %s`
		case reflect.Float32, reflect.Float64,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			op.neq = `%s != %s`
		case reflect.String:
			op.neq = `%s != %s`
		case reflect.Struct:
			op.neq = `!%s.Equal(%s)`
		}
	}

	if op.nzero == "" {
		switch op.kind {
		case reflect.Bool:
			op.nzero = `%s`
		case reflect.Float32, reflect.Float64,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			op.nzero = `%s != 0`
		case reflect.String:
			op.nzero = `%s != ""`
		case reflect.Struct:
			op.nzero = `!%s.IsZero()`
		case reflect.Slice:
			op.nzero = `%s != nil`
		}
	}

	if op.assign == "" {
		return fmt.Errorf("assign is empty (kind=%s)", op.kind)
	}

	if op.neq == "" {
		return fmt.Errorf("neq is empty (kind=%s)", op.kind)
	}

	if op.nzero == "" {
		return fmt.Errorf("nzero is empty (kind=%s)", op.kind)
	}

	return nil
}

func (op *typeOp) NotZero(v string) string {
	return fmt.Sprintf(op.nzero, v)
}

func (op *typeOp) NotEqual(left, right string) string {
	return fmt.Sprintf(op.neq, left, right)
}

func (op *typeOp) Assign(left, right string) string {
	return fmt.Sprintf(op.assign, left, right)
}
