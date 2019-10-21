/*
Q21. (1) Cat
1. 编写一个程序模仿 Unix 的 cat 程序。对于不知道这个程序的人来说，下面的
调用显示了文件 blah 的内容：
% cat blah
2. 使其支持 n 开关，用于输出每行的行号。
3. 上面问题中，1 提供的解决方案存在一个 Bug。你能定位并修复它吗？
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var numberFlag = flag.Bool("n", false, "number each line")

func cat(r *bufio.Reader) {
	i := 1
	for {
		buf, e := r.ReadBytes('\n')
		if e == io.EOF {
			break
		}
		if *numberFlag {
			fmt.Fprintf(os.Stdout, "%5d %s", i, buf)
			i++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, e := os.Open(flag.Arg(i))
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n", os.Args[0], flag.Arg(i), e.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}

/*
这个代码和之前的一样的，书上例子估计有问题
2. 当最后一行不包括换行符时，这个 Bug 就会出现。更糟糕的情况是，当输入只
有一行且没有换行符的时候，什么也不显示。下面的程序是一个更好的解决方案。
*/
