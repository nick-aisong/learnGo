// 6-2 hash2.go
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	TestFile := "hash2.go"
	infile, inerr := os.Open(TestFile)
	if inerr == nil {
		md5h := md5.New()
		io.Copy(md5h, infile)
		fmt.Printf("%x %s\n", md5h.Sum([]byte("")), TestFile)
		sha1h := sha1.New()
		io.Copy(sha1h, infile)
		fmt.Printf("%x %s\n", sha1h.Sum([]byte("")), TestFile)
	} else {
		fmt.Println(inerr)
		os.Exit(1)
	}
}

//G:\GitHub\learnGo\Go语言编程\ch06_安全编程>go run hash2.go
//3fc83c87fe7eace05cd260f5488408b5 hash2.go
//da39a3ee5e6b4b0d3255bfef95601890afd80709 hash2.go
