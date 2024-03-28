package main

import (
	"fmt"
	"go/ast"
)

const Pi = 3.14

var version = "1.0"

type T struct {
	X int
	Y int
}

func (t *T) Info() string {
	return fmt.Sprintf("%v-%v", t.X, t.Y)
}

var box *ast.ArrayType

func main() {
	fmt.Println(version)
	a := 100
	_ = a
}
