package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"gop-learn/gop-learn/2-types/load"

	"golang.org/x/tools/go/ast/astutil"
)

func NodeGo() {
	fset := token.NewFileSet()
	_, info, files, err := load.LoadPackage(fset, "main", "./testdata/test.go")

	// f, err := parser.ParseFile(fset, "./testdata/test.go", nil, parser.ParseComments)
	f := files[0]
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
	fmt.Println(f.Name)
	// paths, exact := astutil.PathEnclosingInterval(f, 121, 121) //t
	paths, exact := astutil.PathEnclosingInterval(f, 189, 189) //box
	fmt.Println(paths, exact)
	if exact {
		ident, ok := paths[0].(*ast.Ident)
		if !ok {
			fmt.Println("ident undefined", info)
			return
		}
		obj := ident.Obj
		_ = obj
		switch obj.Kind {
		case ast.Var:
			fmt.Printf("标识符%s ast类型%s", ident.Name, "var")
			if obj.Decl != nil {
				switch obj.Decl.(type) {
				case *ast.Field:
					if field, ok := obj.Decl.(*ast.Field); ok {
						// obj.Decl 断言为 *ast.Field 类型
						starExpr, ok := field.Type.(*ast.StarExpr)
						if ok {
							// 处理 starExpr
							// ...
							_ = starExpr
							id, okk := starExpr.X.(*ast.Ident)
							if okk {
								fmt.Printf("实际类型 *%s", id.Name)
							}
							// starExpr.X.Name
						}
					}

				default:
					fmt.Println("test")
				}
			}

		default:
			fmt.Println("Unknown type")
		}

	}
}
