/*
Q33. (2) 数字游戏
• 从列表中随机选择六个数字：
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 25, 50, 75, 100
数字可以多次被选中；
• 从 1 . . . 1000 中选择一个随机数 i；
• 尝试用先前的六个数字（或者其中的几个）配合运算符 +，−，∗ 和 /，计算出 i；
例如，选择了数字：1，6，7，8，8 和 75。并且 i 为 977。可以用许多方法来实现，其
中一种：
	((((1 ∗ 6) ∗ 8) + 75) ∗ 8) − 7 = 977
或者
	(8 ∗ (75 + (8 ∗ 6))) − (7/1) = 977
1. 实现像这样的数字游戏。使其打印像上面那样格式的结果（也就是说，输出应
当是带有括号的中序表达式）
2. 计算全部可能解，并且全部显示出来（或者仅显示有多少个）。在上面的例子
中，有 544 种方法。
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
