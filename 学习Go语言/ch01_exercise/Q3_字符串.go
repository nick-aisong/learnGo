/*
Q3. (1) 字符串
1. 建立一个 Go 程序打印下面的内容（到 100 个字符）：
A
AA
AAA
AAAA
AAAAA
AAAAAA
AAAAAAA
...
2. 建立一个程序统计字符串里的字符数量：
asSASA ddd dsjkdsjs dk
同时输出这个字符串的字节数。提示： 看看 unicode/utf8 包
3. 扩展/修改上一个问题的程序，替换位置 4 开始的三个字符为 "abc"
4. 编写一个 Go 程序可以逆转字符串，例如 "foobar" 被打印成 "raboof"
*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func Q1_func1(counter int) {
	for i := 1; i <= counter; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("A")
		}
		fmt.Printf("\n")
	}
}

func Q1_func2(counter int) {
	str := "A"
	for i := 0; i < counter; i++ {
		fmt.Printf("%s\n", str)
		str = str + "A"
	}
}

func Q2_func1(str string) (counter int) {
	counter = 0
	for range str {
		counter++
	}
	return
}

func Q2_func2(str string) {
	fmt.Printf("String: %s\nLength: %d, Runes: %d\n", str, len([]byte(str)), utf8.RuneCount([]byte(str)))
}

func Q3_func1(str string) {
	r := []rune(str)
	copy(r[4:4+3], []rune("abc"))
	fmt.Printf("Before: %s\n", str)
	fmt.Printf("After : %s\n", string(r))
}

func Q4_func1(str string) {
	a := []rune(str)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Printf("Before: %s\n", str)
	fmt.Printf("After : %s\n", string(a))
}

func main() {
	Q1_func1(100)
	Q1_func2(100)

	str := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("%d\n", Q2_func1(str))
	Q2_func2(str)

	Q3_func1(str)

	Q4_func1(str)
}
