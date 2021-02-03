package main

import (
	"fmt"
	"os"
)

func main() {
	/*基础类型变量*/
	var i int
	var f32 float32
	var f64 float64
	var b bool
	var s string
	/*派生类型变量*/
	var file os.File
	fmt.Printf("%v %v %v %v %q %v \n\n", i, f32, f64, b, s, file)
}
