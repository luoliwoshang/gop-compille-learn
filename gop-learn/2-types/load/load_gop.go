package load

import (
	"path/filepath"

	goast "go/ast"
	goparser "go/parser"

	"go/importer"
	"go/types"

	"github.com/goplus/gop/ast"
	"github.com/goplus/gop/parser"
	"github.com/goplus/gop/token"
	"github.com/goplus/gop/x/typesutil"
	"github.com/goplus/mod/gopmod"
)

func LoadGopPackage(fset *token.FileSet, name string, filenames ...string) (*types.Package, *typesutil.Info, error) {
	var files []*ast.File
	var gofiles []*goast.File
	for _, filename := range filenames {
		switch ext := filepath.Ext(filename); ext {
		case ".go":
			f, err := goparser.ParseFile(fset, filename, nil, goparser.ParseComments)
			if err != nil {
				return nil, nil, err
			}
			gofiles = append(gofiles, f)
		case ".gop", ".gox":
			mode := parser.ParseComments
			if ext == ".gox" {
				mode |= parser.ParseGoPlusClass
			}
			f, err := parser.ParseFile(fset, filename, nil, mode)
			if err != nil {
				return nil, nil, err
			}
			files = append(files, f)
		}
	}

	pkg := types.NewPackage("", name)
	conf := &types.Config{}
	conf.Importer = importer.Default()
	chkOpts := &typesutil.Config{
		Types: pkg,
		Fset:  fset,
		Mod:   gopmod.Default,
	}
	info := &typesutil.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Implicits:  make(map[ast.Node]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Scopes:     make(map[ast.Node]*types.Scope),
		Overloads:  make(map[*ast.Ident][]types.Object),
	}
	check := typesutil.NewChecker(conf, chkOpts, nil, info)
	err := check.Files(gofiles, files)
	return pkg, info, err
}
