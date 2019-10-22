/*
Q29. (0) 单词和字母统计
1. 编写一个从标准输入中读取文本的小程序，并进行下面的操作：
	1. 计算字符数量（包括空格）；
	2. 计算单词数量；
	3. 计算行数。
换句话说，实现一个 wc，然而只需要从标准输入读取。
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var chars, words, lines int
	r := bufio.NewReader(os.Stdin)
	for {
		switch s, ok := r.ReadString('\n'); true {
		case ok != nil:
			fmt.Printf("%d %d %d\n", chars, words, lines)
			return
		default:
			chars += len(s)
			words += len(strings.Fields(s))
			lines++
		}
	}
}
