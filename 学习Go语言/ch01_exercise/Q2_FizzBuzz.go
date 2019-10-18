/*
Q2. (0) FizzBuzz
1. 解决这个叫做 Fizz-Buzz 的问题：
编写一个程序，打印从 1 到 100 的数字。
当是三个数字就打印 "Fizz" 代替数字，当是五的倍数就打印 "Buzz"。
当数字同时是三和五的倍数时，打印 "FizzBuzz"。
*/

package main

import (
	"fmt"
)

func FizzBuzz(counter int) {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func FizzBuzz2(counter int) {
	const (
		FIZZ = 3
		BUZZ = 5
	)
	var p bool
	for i := 1; i <= 100; i++ {
		p = false
		if i%FIZZ == 0 {
			fmt.Printf("Fizz")
			p = true
		}
		if i%BUZZ == 0 {
			fmt.Printf("Buzz")
			p = true
		}
		if !p {
			fmt.Printf("%v", i)
		}
		fmt.Println()
	}
}

func main() {
	FizzBuzz(100)
	FizzBuzz2(100)
}
