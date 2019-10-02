// once
package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var a string
var once sync.Once

func setup() {
	a = "hello, world"
	count = count + 1
}

func doprint() {
	once.Do(setup)
	fmt.Println(a)
}

func twoprint() {
	go doprint()
	go doprint()
	go doprint()
	go doprint()
	go doprint()
	go doprint()
}

func main() {
	twoprint()
	time.Sleep(time.Second)
	fmt.Println(count)
}
