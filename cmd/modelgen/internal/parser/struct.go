package modelgen

import (
	"go/token"
	"go/types"
	"reflect"
	"strings"
)

const (
	TagDB    = "db"
	TagModel = "model"
)

type Struct struct {
	Name   string
	N      *types.Named
	S      *types.Struct
	Fields []*StructField

	p             *Package
	fieldSet      []*StructField
	dbFieldSet    []*StructField
	modelFieldSet []*StructField
}

// NewStruct creates a new Struct.
// It panics if the underlying type of `named` is not `*types.Struct`.
func NewStruct(p *Package, named *types.Named) *Struct {
	ts := named.Underlying().(*types.Struct)

	s := &Struct{
		Name:   named.Obj().Name(),
		S:      ts,
		N:      named,
		Fields: make([]*StructField, ts.NumFields()),
		p:      p,
	}

	for i := 0; i < ts.NumFields(); i++ {
		v := ts.Field(i)
		tag := ts.Tag(i)
		s.Fields[i] = NewStructField(v, tag)
	}

	return s
}

func (s *Struct) Pos() token.Pos {
	return s.N.Obj().Pos()
}

func (s *Struct) Position() token.Position {
	return s.p.Fset.Position(s.Pos())
}

func (s *Struct) FindSibling(name string) *Struct {
	return s.p.Structs[name]
}

func (s *Struct) FindSiblingByType(t types.Type) *Struct {
	if sl, ok := t.(*types.Slice); ok {
		t = sl.Elem()
	}

	if ptr, ok := t.(*types.Pointer); ok {
		t = ptr.Elem()
	}

	if named, ok := t.(*types.Named); ok {
		return s.FindSibling(named.Obj().Name())
	}

	return nil
}

func (s *Struct) Field(name string) *StructField {
	for _, f := range s.Fields {
		if f.Name() == name {
			return f
		}
	}

	return nil
}

// FieldSet collects all fields, including fields of embedded types.
func (s *Struct) FieldSet() []*StructField {
	if s.fieldSet != nil {
		return s.fieldSet
	}

	s.fieldSet = make([]*StructField, 0, len(s.Fields))

	for _, f := range s.Fields {
		if f.Embedded() {
			if named, ok := f.Type().(*types.Named); ok {
				if _, ok := named.Underlying().(*types.Struct); ok {
					var pkg *Package

					if named.Obj().Pkg().Path() == s.N.Obj().Pkg().Path() { // from same package
						pkg = s.p
					} else { // from imported package
						pkg = s.p.FindImport(named.Obj())
					}

					es := pkg.Structs[named.Obj().Name()] // panics with pkg == nil
					fs := es.FieldSet()                   // panics with es == nil
					s.fieldSet = append(s.fieldSet, fs...)

					continue
				}
			}
		}

		s.fieldSet = append(s.fieldSet, f)
	}

	return s.fieldSet
}

// DBFieldSet collects all fields, including fields of embedded types, that have a valid `TagDB` tag.
func (s *Struct) DBFieldSet() []*StructField {
	if s.dbFieldSet != nil {
		return s.dbFieldSet
	}

	s.dbFieldSet = make([]*StructField, 0, len(s.FieldSet()))

	for _, f := range s.FieldSet() {
		if f.IsDB() {
			s.dbFieldSet = append(s.dbFieldSet, f)
		}
	}

	return s.dbFieldSet
}

// DBFieldColMap maps field names from `DBFieldSet` to their respective database columns.
func (s *Struct) DBFieldColMap() map[string]string {
	dbFields := s.DBFieldSet()
	fieldCols := make(map[string]string, len(dbFields))

	for _, f := range dbFields {
		fieldCols[f.Name()] = f.DBStructTag().Col
	}

	return fieldCols
}

// ModelFieldSet collects all fields, including fields of embedded types, that have a valid `TagModel` tag.
func (s *Struct) ModelFieldSet() []*StructField {
	if s.modelFieldSet != nil {
		return s.modelFieldSet
	}

	s.modelFieldSet = make([]*StructField, 0, len(s.FieldSet()))

	for _, f := range s.FieldSet() {
		if f.IsModel() {
			s.modelFieldSet = append(s.modelFieldSet, f)
		}
	}

	return s.modelFieldSet
}

// IsModelStruct considers a struct as being a model if it has an ID field with a valid `TagDB` tag.
//
// It doesn't check if the struct implements the `Model` interface because the user may be trying
// generate code that specifically satisfy the interface.
func (s *Struct) IsModelStruct() bool {
	f := s.Field("ID")
	return f != nil && f.IsDB()
}

func (s *Struct) GetFirstRecvName(def string) string {
	for i := 0; i < s.N.NumMethods(); i++ {
		fn := s.N.Method(i)
		sig := fn.Type().(*types.Signature)

		if n := sig.Recv().Name(); n != "" {
			return n
		}
	}

	return def
}

type StructField struct {
	*types.Var

	Tag reflect.StructTag

	dbTag    *DBStructTag
	modelTag *ModelStructTag
}

func NewStructField(v *types.Var, tag string) *StructField {
	return &StructField{
		Var: v,
		Tag: reflect.StructTag(tag),
	}
}

func (f *StructField) DBStructTag() *DBStructTag {
	if f.dbTag != nil {
		return f.dbTag
	}

	f.dbTag = NewDBStructTag(f.Tag)

	return f.dbTag
}

func (f *StructField) IsDB() bool {
	if !f.Exported() {
		return false
	}

	t := f.DBStructTag()

	return t != nil && !t.Ignore
}

func (f *StructField) ModelStructTag() *ModelStructTag {
	if f.modelTag != nil {
		return f.modelTag
	}

	f.modelTag = NewModelStructTag(f.Tag)

	return f.modelTag
}

func (f *StructField) IsModel() bool {
	if !f.Exported() {
		return false
	}

	t := f.ModelStructTag()

	return t != nil && !t.Ignore
}

type DBStructTag struct {
	Ignore bool
	Col    string
}

// NewDBStructTag parses a `db:"<content>"` struct tag.
//
// Content format: "-"|(column_name[(","option)...])
func NewDBStructTag(tag reflect.StructTag) *DBStructTag {
	tagVal := tag.Get(TagDB)

	if tagVal == "" {
		return nil
	}

	dbTag := &DBStructTag{}

	if tagVal == "-" {
		dbTag.Ignore = true
		return dbTag
	}

	opts := strings.Split(tagVal, ",")

	dbTag.Col = opts[0]

	for _, opt := range opts[1:] {
		switch opt {
		// no-op
		}
	}

	return dbTag
}

type ModelStructTag struct {
	Ignore bool
	Type   string
	Params map[string]string
}

// NewModelStructTag parses a `model:"<content>"` struct tag.
//
// Content format: "-"|(association_type[(","key["="value])...])
func NewModelStructTag(tag reflect.StructTag) *ModelStructTag {
	tagVal := tag.Get(TagModel)

	if tagVal == "" {
		return nil
	}

	modelTag := &ModelStructTag{}

	if tagVal == "-" {
		modelTag.Ignore = true
		return modelTag
	}

	pairs := strings.Split(tagVal, ",")

	modelTag.Type = pairs[0]
	modelTag.Params = make(map[string]string)

	for _, pair := range pairs[1:] {
		var val string

		kv := strings.SplitN(pair, ":", 2)
		key := kv[0]

		if len(kv) > 1 {
			val = kv[1]
		}

		modelTag.Params[key] = val
	}

	return modelTag
}
