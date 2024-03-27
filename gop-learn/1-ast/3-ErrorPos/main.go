package main

import (
	"fmt"
	"go/parser"
	"go/scanner"
	"go/token"
)

func main() {
	// 创建一个新的FileSet
	fset := token.NewFileSet()

	// 解析源代码文件
	file, err := parser.ParseFile(fset, "./testdata/test.go", nil, parser.ParseComments)

	// 缺少闭括号，引发语法错误
	if err != nil {
		// for scErr:=range err
		scannerErrors := err.(scanner.ErrorList)
		for _, scannerError := range scannerErrors {
			// 处理每个 scannerError
			fmt.Println("Syntax error:", scannerError.Pos)
		}
		return
	}
	_ = file
	// // 获取语法错误信息
	// for _, err := range file.Comments {
	// 	for _, comment := range err.List {
	// 		pos := comment.Pos()
	// 		position := fset.Position(pos)

	// 		fmt.Printf("Syntax error at line %d, column %d: %s\n",
	// 			position.Line, position.Column, comment.Text)
	// 	}
	// }
}
