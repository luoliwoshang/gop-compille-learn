package main

import (
	"fmt"
	"go/token"
	"go/types"
	"gop-learn/gop-learn/2-types/load"
)

func main() {
	fset := token.NewFileSet()
	pkg, info, err := load.LoadPackage(fset, "main", "./testdata/test.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(pkg)
	fmt.Println(pkg.Scope())
	fmt.Println(info)
	// 在类型检查的结果中获取信息
	for ident, obj := range info.Defs {
		// 只检查变量定义
		if v, ok := obj.(*types.Var); ok {
			// 检查变量是否被使用
			fmt.Printf("标识符 %s 的对象类型为: %s\n", ident.Name, v.Type())
			name := ident.Name
			used := false
			// Uses中的标识符存储的位置是使用的位置，而Defs存储的位置是定义的位置，所以通过ident.name来判断
			for _ident := range info.Uses {
				if _ident.Name == name {
					used = true
				}
			}
			if !used {
				fmt.Printf("Variable %s is unused\n", ident.Name)
			}
		}

	}
}
