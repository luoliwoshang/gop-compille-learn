package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "./testdata/test.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
	// 从ast树中深度优先遍历每一个节点（Node）
	ast.Inspect(f, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.CallExpr:
			fmt.Printf("%#v %v\n", n, fset.Position(n.Pos()))
			// &ast.CallExpr{Fun:(*ast.SelectorExpr)(0xc000010018), Lparen:96, Args:[]ast.Expr{(*ast.BasicLit)(0xc000028220)}, Ellipsis:0, Rparen:110} ./testdata/test.go:11:2
		}
		return true
	})
}
