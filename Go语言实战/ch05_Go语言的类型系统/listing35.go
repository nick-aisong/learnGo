// 5-35 listing35.go
// 这个示例程序展示 bytes.Buffer 也可以
// 用于 io.Copy 函数
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// main 是应用程序的入口
func main() {
	var b bytes.Buffer

	// 将字符串写入 Buffer
	b.Write([]byte("Hello"))

	// 使用 Fprintf 将字符串拼接到 Buffer
	fmt.Fprintf(&b, "World!")

	// 将 Buffer 的内容写到 Stdout
	io.Copy(os.Stdout, &b)
}
