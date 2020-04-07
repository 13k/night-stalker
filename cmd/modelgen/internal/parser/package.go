package modelgen

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/packages"
)

type Package struct {
	*packages.Package

	T *types.Package
	S *types.Scope

	// Collected named types

	Structs    map[string]*Struct
	Interfaces map[string]*Interface
	Functions  map[string]*types.Func
	Types      map[string]*types.Named

	ModelInterface *Interface
}

func NewPackage(pkg *packages.Package) *Package {
	p := &Package{
		Package:    pkg,
		T:          pkg.Types,
		Structs:    make(map[string]*Struct),
		Interfaces: make(map[string]*Interface),
		Functions:  make(map[string]*types.Func),
		Types:      make(map[string]*types.Named),
	}

	if p.T != nil {
		p.S = p.T.Scope()
		p.collectNamedTypes()
	}

	return p
}

func (p *Package) collectNamedTypes() {
	for _, name := range p.S.Names() {
		obj := p.S.Lookup(name)

		if obj == nil {
			continue
		}

		if !obj.Exported() {
			continue
		}

		switch t := obj.(type) {
		case *types.Func:
			p.Functions[t.Name()] = t
		case *types.TypeName:
			switch ot := t.Type().(type) {
			case *types.Named:
				switch ot.Underlying().(type) {
				case *types.Struct:
					p.Structs[t.Name()] = NewStruct(p, ot)
				case *types.Interface:
					p.Interfaces[t.Name()] = NewInterface(p, ot)
				default:
					p.Types[t.Name()] = ot
				}
			default:
				// TypeName with unknown Type
			}
		default:
			// Unknown Object
		}
	}

	p.ModelInterface = p.Interfaces["Model"]
}

func (p *Package) ImplementsModel(t types.Type) bool {
	if p.ModelInterface == nil {
		return false
	}

	if p.ModelInterface.ImplementedBy(t) {
		return true
	}

	// try pointer to struct
	if _, ok := t.Underlying().(*types.Struct); ok {
		t = types.NewPointer(t)

		if p.ModelInterface.ImplementedBy(t) {
			return true
		}
	}

	return false
}

func (p *Package) FindImport(obj types.Object) *Package {
	ppkg := p.Imports[obj.Pkg().Path()]

	if ppkg == nil {
		fmt.Printf("FindImport Query=%q\n", obj.Pkg().Path())

		for p := range p.Imports {
			fmt.Printf("  Available Path=%q\n", p)
		}

		return nil
	}

	return NewPackage(ppkg)
}
