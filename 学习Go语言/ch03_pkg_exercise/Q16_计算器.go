/*
1. 使用 stack 包创建逆波兰计算器
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stack struct {
	i    int
	data [10]int
}

func (s *Stack) push(e int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = e
	s.i++
}

func (s *Stack) pop() (e int) {
	s.i--
	if s.i < 0 {
		s.i = 0
		return
	}
	return s.data[s.i]
}

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var st = new(Stack)

func main() {
	for {
		s, err := reader.ReadString('\n')
		var token string
		if err != nil {
			return
		}
		for _, c := range s {
			switch {
			case c >= '0' && c <= '9':
				token = token + string(c)
			case c == ' ':
				r, _ := strconv.Atoi(token)
				st.push(r)
				token = ""
			case c == '+':
				fmt.Printf("%d\n", st.pop()+st.pop())
			case c == '*':
				fmt.Printf("%d\n", st.pop()*st.pop())
			case c == '-':
				p := st.pop()
				q := st.pop()
				fmt.Printf("%d\n", q-p)
			case c == 'q':
				return
			default:
				// error
			}
		}
	}
}

// a+b       ->  ab+
// a+(b-c)   ->  abc-+
// a+(b-c)*d ->  adbc-*+

// 这个简单示例只能进行两个数四则运算

// G:\GitHub\learnGo\学习Go语言\ch03_pkg_exercise>go run Q16_计算器.go
// 1 1 +
// 2
