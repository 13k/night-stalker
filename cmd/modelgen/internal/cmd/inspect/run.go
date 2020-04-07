package cmdinspect

import (
	"fmt"
	"go/types"

	g "github.com/13k/night-stalker/cmd/modelgen/internal/parser"
)

type Options struct {
	Package string
}

func Run(opts *Options) error {
	pkgs, err := g.Parse(opts.Package)

	if err != nil {
		return fmt.Errorf("error parsing package %s: %w", opts.Package, err)
	}

	if n := g.PrintErrors(pkgs); n > 0 {
		return fmt.Errorf("found %d errors while parsing package %s", n, opts.Package)
	}

	for _, pkg := range pkgs {
		printPkg(pkg)
	}

	return nil
}

func printPkg(p *g.Package) {
	fmt.Printf("Package Name=%q\n", p.Name)

	for _, x := range p.Imports {
		fmt.Printf("  Import Path=%q\n", x.PkgPath)
	}

	for _, s := range p.Structs {
		printStruct(p, s)
	}
}

func printStruct(p *g.Package, s *g.Struct) {
	position := s.Position()

	fmt.Printf(
		"  Struct Name=%q Fields=%d IsModel=%v ImplementsModel=%v (%s:%d)\n",
		s.Name,
		s.S.NumFields(),
		s.IsModelStruct(),
		p.ImplementsModel(s.N),
		position.Filename,
		position.Line,
	)

	for _, f := range s.Fields {
		fmt.Printf(
			"    Field Name=%q Embedded=%v Exported=%v Type=%s t=%T Tag=%q\n",
			f.Name(),
			f.Embedded(),
			f.Exported(),
			f.Type(),
			f.Type(),
			f.Tag,
		)
	}

	for i := 0; i < s.N.NumMethods(); i++ {
		m := s.N.Method(i)

		fmt.Printf(
			"    Method Name=%q Sig=%s\n",
			m.Name(),
			m.Type().(*types.Signature).String(),
		)
	}
}
