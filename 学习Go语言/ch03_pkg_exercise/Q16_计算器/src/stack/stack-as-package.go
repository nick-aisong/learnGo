/*
1. 参考 Q8 练习。在这个练习中将从那个代码中建立一个独立的包
为 stack 的实现创建一个合适的包，Push、Pop 和 Stack 类型需要被导出
2. 为这个包编写一个单元测试，至少测试 Push 后 Pop 的工作情况
*/
package stack

// 保存元素的 Stack
type Stack struct {
	i    int
	data [10]int
}

// Push 将元素压入栈中
func (s *Stack) Push(k int) {
	s.data[s.i] = k
	s.i++
}

// Pop 从栈中弹出一个元素
func (s *Stack) Pop() (ret int) {
	s.i--
	ret = s.data[s.i]
	return
}
