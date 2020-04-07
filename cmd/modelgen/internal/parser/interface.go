package modelgen

import (
	"go/token"
	"go/types"
)

type Interface struct {
	Name string
	N    *types.Named
	I    *types.Interface
	p    *Package
}

// NewInterface creates a new Interface.
// It panics if the underlying type of `named` is not `*types.Interface`.
func NewInterface(p *Package, named *types.Named) *Interface {
	return &Interface{
		Name: named.Obj().Name(),
		N:    named,
		I:    named.Underlying().(*types.Interface),
	}
}

func (i *Interface) ImplementedBy(t types.Type) bool {
	return types.Implements(t, i.I)
}

func (i *Interface) Pos() token.Pos {
	return i.N.Obj().Pos()
}

func (i *Interface) Position() token.Position {
	return i.p.Fset.Position(i.Pos())
}
