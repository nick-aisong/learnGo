// arrayDemo
package main

func main() {

	// 4-1 声明一个数组，并设置为零值
	// 声明一个包含 5 个元素的整型数组
	var array [5]int

	// 4-2 使用数组字面量声明数组
	// 声明一个包含 5 个元素的整型数组
	// 用具体值初始化每个元素
	array := [5]int{10, 20, 30, 40, 50}

	// 4-3 让 Go 自动计算声明数组的长度
	// 声明一个整型数组
	// 用具体值初始化每个元素// 容量由初始化值的数量决定
	array := [...]int{10, 20, 30, 40, 50}

	// 4-4 声明数组并指定特定元素的值
	// 声明一个有 5 个元素的数组
	// 用具体值初始化索引为 1 和 2 的元素
	// 其余元素保持零值
	array := [5]int{1: 10, 2: 20}

	// 4-5 访问数组元素
	// 声明一个包含 5 个元素的整型数组
	// 用具体值初始为每个元素
	array := [5]int{10, 20, 30, 40, 50}
	// 修改索引为 2 的元素的值
	array[2] = 35

	// 4-6 访问指针数组的元素
	// 声明包含 5 个元素的指向整数的数组
	// 用整型指针初始化索引为 0 和 1 的数组元素
	array := [5]*int{0: new(int), 1: new(int)}
	// 为索引为 0 和 1 的元素赋值
	*array[0] = 10
	*array[1] = 20

	// 4-7 把同样类型的一个数组赋值给另外一个数组
	// 声明第一个包含 5 个元素的字符串数组
	var array1 [5]string
	// 声明第二个包含 5 个元素的字符串数组
	// 用颜色初始化数组
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	// 把 aarray2 的值复制到 aarray1 （值拷贝）
	array1 = array2

	// 4-8 编译器会阻止类型不同的数组互相赋值
	// 声明第一个包含 4 个元素的字符串数组
	var array1 [4]string
	// 声明第二个包含 5 个元素的字符串数组
	// 使用颜色初始化数组
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	// 将 array2 复制给 array1
	array1 = array2
	//Compiler Error:
	//cannot use array2 (type [5]string) as type [4]string in assignment

	// 4-9 把一个指针数组赋值给另一个
	// 声明第一个包含 3 个元素的指向字符串的指针数组
	var array1 [3]*string
	// 声明第二个包含 3 个元素的指向字符串的指针数组
	// 使用字符串指针初始化这个数组
	array2 := [3]*string{new(string), new(string), new(string)}
	// 使用颜色为每个元素赋值
	*array2[0] = "Red"
	*array2[1] = "Blue"
	*array2[2] = "Green"
	// 将 array2 复制给 array1
	array1 = array2

}
