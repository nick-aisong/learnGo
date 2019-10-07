// dictDemo
package main

import (
	"fmt"
)

func main() {
	// 4-45 使用 make 声明映射
	// 创建一个映射，键的类型是 string，值的类型是 int
	dict := make(map[string]int)

	// 创建一个映射，键和值的类型都是 string
	// 使用两个键值对初始化映射
	dict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

	// 4-46 使用映射字面量声明空映射
	// 创建一个映射，使用字符串切片作为映射的键
	dict := map[[]string]int{}
	//Compiler Exception:
	//invalid map key type []string

	// 4-47 声明一个存储字符串切片的映射
	// 创建一个映射，使用字符串切片作为值
	dict := map[int][]string{}

	// 4-48 为映射赋值
	// 创建一个空映射，用来存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{}
	// 将 Red 的代码加入到映射
	colors["Red"] = "#da1337"

	// 4-49 对 nil 映射赋值时的语言运行时错误
	// 通过声明映射创建一个 nil 映射
	var colors map[string]string
	// 将 Red 的代码加入到映射
	colors["Red"] = "#da1337"
	//Runtime Error:
	//panic: runtime error: assignment to entry in nil map

	// 4-50 从映射获取值并判断键是否存在
	// 获取键 Blue 对应的值
	value, exists := colors["Blue"] // 这个键存在吗？
	if exists {
		fmt.Println(value)
	}

	// 4-51 从映射获取值，并通过该值判断键是否存在
	// 获取键 Blue 对应的值
	value := colors["Blue"]
	// 这个键存在吗？
	if value != "" {
		fmt.Println(value)
	}
	// 这种方法只能用在映射存储的值都是非零值的情况
	// 键不存在返回的是该值对应的类型的零值

	// 4-52 使用 range 迭代映射
	// 创建一个映射，存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	// 4-53 从映射中删除一项
	// 删除键为 Coral 的键值对

	delete(colors, "Coral")
	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}
