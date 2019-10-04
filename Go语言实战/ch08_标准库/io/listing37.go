// 8-37 listing37.go
// 这个示例程序展示来自不同标准库的不同函数是如何
// 使用 io.Writer 接口的
package main

import (
	"bytes"
	"fmt"
	"os"
)

// main 是应用程序的入口
func main() {
	// 创建一个 Buffer 值，并将一个字符串写入 Buffer
	// 使用实现 io.Writer 的 Write 方法
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	// 使用 Fprintf 来将一个字符串拼接到 Buffer 里
	// 将 bytes.Buffer 的地址作为 io.Writer 类型值传入
	fmt.Fprintf(&b, "World!")

	// 将 Buffer 的内容输出到标准输出设备
	// 将 os.File 值的地址作为 io.Writer 类型值传入
	b.WriteTo(os.Stdout)
}

// 8-39 golang.org/src/fmt/print.go
// Fprintf 根据格式化说明符来格式写入内容，并输出到 w
// 这个函数返回写入的字节数，以及任何遇到的错误
// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)

// 8-40 golang.org/src/bytes/buffer.go
// Write 将 p 的内容追加到缓冲区，如果需要，会增大缓冲区的空间。返回值 n 是
// p 的长度， err 总是 nil。如果缓冲区变得太大， Write 会引起崩溃…
// func (b *Buffer) Write(p []byte) (n int, err error) {
// 	b.lastRead = opInvalid
// 	m := b.grow(len(p))
// 	return copy(b.buf[m:], p), nil
// }

// 8-42 golang.org/src/os/file.go
// var (
// 	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
// 	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
// 	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
// )

// 8-43 golang.org/src/os/file_unix.go
// NewFile 返回一个具有给定的文件描述符和名字的新 File
// func NewFile(fd uintptr, name string) *File {
// 	fdi := int(fd)
// 	if fdi < 0 {
// 		return nil
// 	}
// 	f := &File{&file{fd: fdi, name: name}}
// 	runtime.SetFinalizer(f.file, (*file).close)
// 	return f
// }

// 8-44 golang.org/src/os/file.go
// Write 将 len(b)个字节写入 File
// 这个方法返回写入的字节数，如果有错误，也会返回错误
// 如果 n != len(b)， Write 会返回一个非 nil 的错误
// func (f *File) Write(b []byte) (n int, err error) {
// 	if f == nil {
// 		return 0, ErrInvalid
// 	}
// 	n, e := f.write(b)
// 	if n < 0 {
// 		n = 0
// 	}
// 	if n != len(b) {
// 		err = io.ErrShortWrite
// 	}
// 	epipecheck(f, e)
// 	if e != nil {
// 		err = &PathError{"write", f.name, e}
// 	}
// 	return n, err
// }
