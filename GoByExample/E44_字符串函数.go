// E44_字符串函数
// 标准库的 strings 包提供了很多有用的字符串相关的函数
// 这里是一些用来让你对这个包有个初步了解的例子
package main

import (
	"fmt"
	s "strings"
)

// 我们给 fmt.Println 一个短名字的别名，我们随后将会经常用到
var p = fmt.Println

func main() {
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	// 你可以在 strings包文档中找到更多的函数
	p()
	// 虽然不是 strings 的一部分，但是仍然值得一提的是获取字符串长度和通过索引获取一个字符的机制
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
