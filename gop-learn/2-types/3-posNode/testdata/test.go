package testdata

import "fmt"

var a = 1
var b = 2

func Hello() string {
	c := a
	_ = c
	fmt.Println("hello world")
	return "hhhh"
}
