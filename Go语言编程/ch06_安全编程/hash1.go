// 6-1 hash1.go
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	TestString := "Hi,pandaman!"

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(TestString))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)

	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestString))
	Result = Sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)
}

//G:\GitHub\learnGo\Go语言编程\ch06_安全编程>go run hash1.go
//70af690adf124b2e828b4d8904ad142a

//49d37c25ef2a62b4cfcf3689c649510390e14875
