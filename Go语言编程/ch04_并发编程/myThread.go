// myThread
package main

import (
	"fmt"
	//	"runtime"
	//	"sync"
)

var counter int = 0

func Count() {
	counter++
	fmt.Println(counter)
}
func main() {
	//lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go Count()
	}

	for {
		//		lock.Lock()
		//		c := counter
		//		lock.Unlock()
		//		runtime.Gosched()
		//		if c >= 10 {
		//			break
		//		}
	}
}

//G:\GitHub\learnGo\Go语言编程\ch04_并发编程>go run myThread.go
//2
//10
//4
//5
//1
//7
//8
//6
//9
//3
//exit status 2
