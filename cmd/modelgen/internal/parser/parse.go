package modelgen

import (
	"golang.org/x/tools/go/packages"
)

func Parse(pattern string) ([]*Package, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedTypes,
	}

	pkgs, err := packages.Load(cfg, pattern)

	if err != nil {
		return nil, err
	}

	ppkgs := make([]*Package, len(pkgs))

	for i, pkg := range pkgs {
		ppkgs[i] = NewPackage(pkg)
	}

	return ppkgs, nil
}

func PrintErrors(pkgs []*Package) int {
	ppkgs := make([]*packages.Package, len(pkgs))

	for i, p := range pkgs {
		ppkgs[i] = p.Package
	}

	return packages.PrintErrors(ppkgs)
}
