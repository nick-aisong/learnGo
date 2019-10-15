// 9-2 cgo1
package main

import "fmt"

/*
#include <stdlib.h>
*/
import "C"

func Random() int {
	return int(C.random())
}
func Seed(i int) {
	C.srandom(C.uint(i))
}
func main() {
	Seed(100)
	fmt.Println("Random:", Random())
}

// 这种方法也行
// #include <stdio.h>
// #include <stdlib.h>
// import "C"

// # command-line-arguments
// exec: "gcc": executable file not found in %PATH%
