package load

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
)

func LoadPackage(fset *token.FileSet, name string, filenames ...string) (*types.Package, *types.Info, []*ast.File, error) {
	var files []*ast.File
	for _, filename := range filenames {
		// 解析每个源代码文件

		f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
		if err != nil {
			return nil, nil, files, err
		}
		files = append(files, f)
	}

	// 创建类型检查器的配置
	conf := &types.Config{}
	conf.Importer = importer.Default()

	// 创建一个新的包和类型检查的信息对象
	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Implicits:  make(map[ast.Node]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Scopes:     make(map[ast.Node]*types.Scope),
	}
	pkg := types.NewPackage("", name)

	// 	types.NewChecker 是 go/types 包中的一个函数，用于创建一个类型检查器的实例。
	// conf 是 types.Config 类型的对象，包含了类型检查器的配置信息。
	// fset 是 *token.FileSet 类型的对象，用于跟踪源代码文件的位置信息。
	// pkg 是 *types.Package 类型的对象，表示要进行类型检查的包。
	// info 是 *types.Info 类型的对象，用于存储类型检查的结果信息。
	check := types.NewChecker(conf, fset, pkg, info)
	err := check.Files(files)
	return pkg, info, files, err
}
