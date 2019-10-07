// 8-11 listing11.go
// 这个示例程序展示如何创建定制的日志记录器
package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger // 记录所有日志
	Info    *log.Logger // 重要的信息
	Warning *log.Logger // 需要注意的信息
	Error   *log.Logger // 非常严重的问题
)

func init() {
	file, err := os.OpenFile("errors.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}

// 8-22 不同的日志方法的声明
// func (l *Logger) Fatal(v ...interface{})
// func (l *Logger) Fatalf(format string, v ...interface{})
// func (l *Logger) Fatalln(v ...interface{})
// func (l *Logger) Flags() int
// func (l *Logger) Output(calldepth int, s string) error
// func (l *Logger) Panic(v ...interface{})
// func (l *Logger) Panicf(format string, v ...interface{})
// func (l *Logger) Panicln(v ...interface{})
// func (l *Logger) Prefix() string
// func (l *Logger) Print(v ...interface{})
// func (l *Logger) Printf(format string, v ...interface{})
// func (l *Logger) Println(v ...interface{})
// func (l *Logger) SetFlags(flag int)
// func (l *Logger) SetPrefix(prefix string)

// G:\GitHub\learnGo\Go语言实战\ch08_标准库\log>go run listing11.go
// INFO: 2019/10/07 18:31:24 listing11.go:45: Special Information
// WARNING: 2019/10/07 18:31:24 listing11.go:46: There is something you need to know about
// ERROR: 2019/10/07 18:31:24 listing11.go:47: Something has failed
