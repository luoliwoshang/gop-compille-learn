package main

import (
	"fmt"
	"gop-learn/gop-learn/2-types/load"

	"github.com/goplus/gop/token"

	"golang.org/x/tools/gop/ast/astutil"
)

func main() {
	fset := token.NewFileSet()
	pkg, info, files, err := load.LoadGopPackage(fset, "main", "./testdata/test.gop")
	_, _, _ = pkg, info, err
	if len(files) > 0 {
		paths, exact := astutil.PathEnclosingInterval(files[0], 47, 47)
		fmt.Println(paths, exact)
		for _, path := range paths {
			fmt.Printf("%#v %v\n", path, fset.Position(path.Pos()))
		}
	}

}
