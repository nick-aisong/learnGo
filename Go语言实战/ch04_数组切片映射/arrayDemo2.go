// arrayDemo2
package main

func main() {

	// 4-10 声明二维数组
	// 声明一个二维整型数组，两个维度分别存储 4 个元素和 2 个元素
	var array [4][2]int
	// 使用数组字面量来声明并初始化一个二维整型数组
	array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// 声明并初始化外层数组中索引为 1 个和 3 的元素
	array := [4][2]int{1: {20, 21}, 3: {40, 41}}
	// 声明并初始化外层数组和内层数组的单个元素
	array := [4][2]int{1: {0: 20}, 3: {1: 41}}

	// 4-11 访问二维数组的元素
	// 声明一个 2×2 的二维整型数组
	var array5 [2][2]int
	// 设置每个元素的整型值
	array[0][0] = 10
	array[0][1] = 20
	array[1][0] = 30
	array[1][1] = 40

	// 4-12 同样类型的多维数组赋值
	// 声明两个不同的二维整型数组
	var array1 [2][2]int
	var array2 [2][2]int
	// 为每个元素赋值
	array2[0][0] = 10
	array2[0][1] = 20
	array2[1][0] = 30
	array2[1][1] = 40
	// 将 array2 的值复制给 array1
	array1 = array2

	// 4-13 使用索引为多维数组赋值
	// 将 array1 的索引为 1 的维度复制到一个同类型的新数组里
	var array3 [2]int = array1[1]
	// 将外层数组的索引为 1、内层数组的索引为 0 的整型值复制到新的整型变量里
	var value int = array1[1][0]

}
