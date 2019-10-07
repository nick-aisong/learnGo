// sliceDemo
package main

import (
	"fmt"
)

func main() {
	// 4-16 使用长度声明一个字符串切片
	// 创建一个字符串切片
	// 其长度和容量都是 5 个元素
	slice := make([]string, 5)

	// 4-17 使用长度和容量声明整型切片
	// 创建一个整型切片
	// 其长度为 3 个元素，容量为 5 个元素
	slice := make([]int, 3, 5)

	// 4-18 容量小于长度的切片会在编译时报错
	// 创建一个整型切片
	// 使其长度大于容量
	// slice := make([]int, 5, 3)
	//Compiler Error:
	//len larger than cap in make([]int)

	// 4-19 通过切片字面量来声明切片
	// 创建字符串切片
	// 其长度和容量都是 5 个元素
	slice := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	// 创建一个整型切片
	// 其长度和容量都是 3 个元素
	slice := []int{10, 20, 30}

	// 4-20 使用索引声明切片
	// 创建字符串切片
	// 使用空字符串初始化第 100 个元素
	slice := []string{99: ""}

	// 4-21 声明数组和声明切片的不同
	// 创建有 3 个元素的整型数组
	array := [3]int{10, 20, 30}
	// 创建长度和容量都是 3 的整型切片
	slice := []int{10, 20, 30}

	// 4-22 创建 nil 切片
	// 创建 nil 整型切片
	var slice []int

	// 4-23 声明空切片
	// 使用 make 创建空的整型切片
	slice := make([]int, 0)
	// 使用切片字面量创建空的整型切片
	slice := []int{}

	// 4-24 使用切片字面量来声明切片
	// 创建一个整型切片
	// 其容量和长度都是 5 个元素
	slice := []int{10, 20, 30, 40, 50}
	// 改变索引为 1 的元素的值
	slice[1] = 25

	// 4-25 使用切片创建切片
	// 创建一个整型切片
	// 其长度和容量都是 5 个元素
	slice := []int{10, 20, 30, 40, 50}
	// 创建一个新切片
	// 其长度为 2 个元素，容量为 4 个元素
	newSlice := slice[1:3]

	// 4-28 修改切片内容可能导致的结果
	// 创建一个整型切片
	// 其长度和容量都是 5 个元素
	slice := []int{10, 20, 30, 40, 50}
	// 创建一个新切片
	// 其长度是 2 个元素，容量是 4 个元素
	newSlice := slice[1:3]
	// 修改 newSlice 索引为 1 的元素
	// 同时也修改了原来的 slice 的索引为 2 的元素
	newSlice[1] = 35

	// 4-29 表示索引越界的语言运行时错误
	// 创建一个整型切片
	// 其长度和容量都是 5 个元素
	slice := []int{10, 20, 30, 40, 50}
	// 创建一个新切片
	// 其长度为 2 个元素，容量为 4 个元素
	newSlice := slice[1:3]
	// 修改 newSlice 索引为 3 的元素
	// 这个元素对于 newSlice 来说并不存在
	newSlice[3] = 45
	//Runtime Exception:
	//panic: runtime error: index out of range

	// 4-30 使用 append 向切片增加元素
	// 创建一个整型切片
	// 其长度和容量都是 5 个元素
	slice := []int{10, 20, 30, 40, 50}
	// 创建一个新切片
	// 其长度为 2 个元素，容量为 4 个元素
	newSlice := slice[1:3]
	// 使用原有的容量来分配一个新元素
	// 将新元素赋值为 60
	newSlice = append(newSlice, 60)

	// 4-31 使用 append 同时增加切片的长度和容量
	// 创建一个整型切片
	// 其长度和容量都是 4 个元素
	slice := []int{10, 20, 30, 40} // 向切片追加一个新元素
	// 将新元素赋值为 50
	newSlice := append(slice, 50)
	// 在切片的容量小于 1000 个元素时，总是会成倍地增加容量。一旦元素个数超过 1000
	// 容量的增长因子会设为 1.25，也就是会每次增加 25%的容量

	// 4-32 使用切片字面量声明一个字符串切片
	// 创建字符串切片
	// 其长度和容量都是 5 个元素
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	// 将第三个元素切片，并限制容量
	// 其长度为 1 个元素，容量为 2 个元素
	slice := source[2:3:4]
	// 对于 slice[i:j:k] 或 [2:3:4]
	// 长度: j – i 或 3 - 2 = 1
	// 容量: k – i 或 4 - 2 = 2

	// 这个切片操作试图设置容量为 4
	// 这比可用的容量大
	slice := source[2:3:6]
	//Runtime Error:
	//panic: runtime error: slice bounds out of range

	// 4-36 设置长度和容量一样的好处
	// 创建字符串切片
	// 其长度和容量都是 5 个元素
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// 对第三个元素做切片，并限制容量
	// 其长度和容量都是 1 个元素
	slice := source[2:3:3]
	// 向 slice 追加新字符串
	slice = append(slice, "Kiwi")

	// 4-37 将一个切片追加到另一个切片
	// 创建两个切片，并分别用两个整数进行初始化
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	// 将两个切片追加在一起，并显示结果
	fmt.Printf("%v\n", append(s1, s2...))
	//Output:
	//[1 2 3 4]

	// 4-38 使用 for range 迭代切片
	// 创建一个整型切片
	// 其长度和容量都是 4 个元素
	slice := []int{10, 20, 30, 40} // 迭代每一个元素，并显示其值
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}
	//Output:
	//Index: 0 Value: 10
	//Index: 1 Value: 20
	//Index: 2 Value: 30
	//Index: 3 Value: 40

	// 4-39 range 提供了每个元素的副本
	// 创建一个整型切片
	// 其长度和容量都是 4 个元素
	slice := []int{10, 20, 30, 40}
	// 迭代每个元素，并显示值和地址
	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n",
			value, &value, &slice[index])
	}
	//Output:
	//Value: 10 Value-Addr: 10500168 ElemAddr: 1052E100
	//Value: 20 Value-Addr: 10500168 ElemAddr: 1052E104
	//Value: 30 Value-Addr: 10500168 ElemAddr: 1052E108
	//Value: 40 Value-Addr: 10500168 ElemAddr: 1052E10C

	// 4-40 使用空白标识符（下划线）来忽略索引值
	// 创建一个整型切片
	// 其长度和容量都是 4 个元素
	slice := []int{10, 20, 30, 40}
	// 迭代每个元素，并显示其值
	for _, value := range slice {
		fmt.Printf("Value: %d\n", value)
	}
	//Output:
	//Value: 10
	//Value: 20
	//Value: 30
	//Value: 40

	// 4-41 使用传统的 for 循环对切片进行迭代
	// 创建一个整型切片
	// 其长度和容量都是 4 个元素
	slice := []int{10, 20, 30, 40}
	// 从第三个元素开始迭代每个元素
	for index := 2; index < len(slice); index++ {
		fmt.Printf("Index: %d Value: %d\n", index, slice[index])
	}
	//Output:
	//Index: 2 Value: 30
	//Index: 3 Value: 40

	// 4-42 声明多维切片
	// 创建一个整型切片的切片
	slice := [][]int{{10}, {100, 200}}

	// 4-43 组合切片的切片
	// 创建一个整型切片的切片
	slice := [][]int{{10}, {100, 200}}
	// 为第一个切片追加值为 20 的元素
	slice[0] = append(slice[0], 20)
}
