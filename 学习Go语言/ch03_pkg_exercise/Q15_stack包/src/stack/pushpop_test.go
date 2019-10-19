// pushpop_test
package stack

import "testing"

func TestPushPup(t *testing.T) {
	c := new(Stack)
	c.Push(5)
	if c.Pop() != 5 {
		t.Log("Pop doesn't give 5")
		t.Fail()
	}
}

// G:\GitHub\learnGo\学习Go语言\ch03_pkg_exercise\Q15_stack包\src\stack>set GOPATH=G:\GitHub\learnGo\学习Go语言\ch03_pkg_exercise\Q15_stack包
// G:\GitHub\learnGo\学习Go语言\ch03_pkg_exercise\Q15_stack包\src\stack>go test stack
// ok      stack   0.243s
