/*
Q31. (2) Quine Quine 是一个打印自己的程序。
1. 用 Go 编写一个 Quine 程序。
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%s%c%s%c\n", q, 0x60, q, 0x60)
}

var q = `
/*
Q31. (2) Quine Quine 是一个打印自己的程序。
1. 用 Go 编写一个 Quine 程序。
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%s%c%s%c\n", q, 0x60)
}
var q
`
