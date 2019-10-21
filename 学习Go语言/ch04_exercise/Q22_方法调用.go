/*
Q22. (2) 方法调用
1. 假设有下面的程序。要注意的是包 container/vector 曾经是 Go 的一部分，但是
当内建的 append 出现后，就被移除了。然而，对于当前的问题这不重要。这个包实现了有 push 和 pop 方法的栈结构。
package main
import "container/vector"
func main() {
	k1 := vector.IntVector{ }
	k2 := &vector.IntVector{ }
	k3 := new(vector.IntVector)
	k1.Push(2)
	k2.Push(3)
	k3.Push(4)
}
k1，k2 和 k3 的类型是什么？
2. 当前，这个程序可以编译并且运行良好。在不同类型的变量上 Push 都可以工
作。Push 的文档这样描述：
func (p *IntVector) Push(x int) Push 增加 x 到向量的末尾。
那么接受者应当是 *IntVector 类型，为什么上面的代码（Push 语句）可以正
确工作？above (the Push statements) work correct then?
*/
/*
1. k1 的类型是 vector.IntVector。为什么？这里使用了符号 {}，因此获得了类型的值。
变量 k2 是 *vector.IntVector，因为获得了复合语句的地址（&）。
而最后的 k3 同样是 *vector.IntVector 类型，因为 new 返回该类型的指针。
2. 在 [10] 的 “调用” 章节，有这样的描述：
	当 x 的方法集合包含 m，并且参数列表可以赋值给 m 的参数，方法调
	用 x.m() 是合法的。如果 x 可以被地址化，而 &x 的方法集合包含 m，
	x.m() 可以作为 (&x).m() 的省略写法。
换句话说，由于 k1 可以被地址化，而 *vector.IntVector 具有 Push 方法，
调用 k1.Push(2) 被 Go 转换为 (&k1).Push(2) 来使型系统愉悦（也使你愉悦
——现在你已经了解到这一点）
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
