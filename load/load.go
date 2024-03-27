package load

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
)

func LoadPackage(fset *token.FileSet, name string, filenames ...string) (*types.Package, *types.Info, error) {
	var files []*ast.File
	for _, filename := range filenames {
		f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
		if err != nil {
			return nil, nil, err
		}
		files = append(files, f)
	}

	conf := &types.Config{}
	conf.Importer = importer.Default()
	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Implicits:  make(map[ast.Node]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Scopes:     make(map[ast.Node]*types.Scope),
	}
	pkg := types.NewPackage("", name)
	check := types.NewChecker(conf, fset, pkg, info)
	err := check.Files(files)
	return pkg, info, err
}
