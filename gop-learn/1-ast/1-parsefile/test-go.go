package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"

	"go/parser"
	"go/token"
)

func TestGo() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "./testdata/test.go", nil, parser.ParseComments)
	// 在这个示例中，使用parser.ParseFile函数解析了名为test.gop的Go源代码文件。
	// 如果源代码中存在语法错误，parser.ParseFile函数将返回一个非空的错误。根据错误信息来定位和修复语法错误。
	if err != nil {
		panic(err)
	}
	for _, i := range f.Imports {
		fmt.Printf("%v %v\n", i.Name, i.Path)
	}
	for _, d := range f.Decls {
		fmt.Printf("%T %v\n", d, d)
	}

	// 获取包名
	packageName := f.Name.Name
	fmt.Println("Package Name:", packageName)

	// 提取函数和方法信息
	for _, decl := range f.Decls {
		// 使用类型断言 decl.(*ast.FuncDecl) 将 decl 转换为 *ast.FuncDecl 类型，然后判断转换是否成功。如果成功，说明该声明是一个函数声明
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			fmt.Println("Function Name:", funcDecl.Name.Name)
			// 可以进一步提取函数参数、返回值等信息
		} else if methodDecl, ok := decl.(*ast.FuncDecl); ok {
			fmt.Println("Method Name:", methodDecl.Name.Name)
			// 可以进一步提取方法接收者、参数、返回值等信息
		}
	}

	// 提取变量信息
	for _, decl := range f.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.VAR {
			for _, spec := range genDecl.Specs {
				if valueSpec, ok := spec.(*ast.ValueSpec); ok {
					for _, ident := range valueSpec.Names {
						fmt.Println("Variable:", ident.Name)
					}
				}
			}
		}
	}

	// 修改函数名
	for _, decl := range f.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			funcDecl.Name.Name = "New" + funcDecl.Name.Name
		}
	}

	// 生成修改后的代码
	var buf bytes.Buffer
	printer.Fprint(&buf, fset, f)
	generatedCode := buf.String()
	fmt.Println(generatedCode)

}
