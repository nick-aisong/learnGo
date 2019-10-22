/*
Q25. (2) 接口和 max()
1. 在练习 Q12 中创建了工作于一个整形 slice 上的最大函数。现在的问题是创建
一个显示最大数字的程序，同时工作于整数和浮点数。虽然在这里会相当困难，
不过还是让程序尽可能的通用吧。
*/
package main

func Less(l, r interface{}) bool {
	switch l.(type) {
	case int:
		if _, ok := r.(int); ok {
			return l.(int) < r.(int)
		}
	case float32:
		if _, ok := r.(float32); ok {
			return l.(float32) < r.(float32)
		}
	}
	return false
}

func main() {
	var a, b, c int = 5, 15, 0
	var x, y, z float32 = 5.4, 29.3, 0.0

	if c = a; Less(a, b) {
		c = b
	}
	if z = x; Less(x, y) {
		z = y
	}
	println(c, z)
}
