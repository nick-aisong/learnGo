// 9-28 listing28_test.go
// 用来检测要将整数值转为字符串，使用哪个函数会更好的基准
// 测试示例。先使用 fmt.Sprintf 函数，然后使用
// strconv.FormatInt 函数，最后使用 strconv.Itoa
package listing28_test

import (
	"fmt"
	"strconv"
	"testing"
)

// BenchmarkSprintf 对 fmt.Sprintf 函数
// 进行基准测试
func BenchmarkSprintf(b *testing.B) {
	number := 10

	//重置计时器，保证测试代码执行前的初始化代码，
	//不会干扰计时器的结果
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}

// BenchmarkFormat 对 strconv.FormatInt 函数
// 进行基准测试
func BenchmarkFormat(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

// BenchmarkItoa 对 strconv.Itoa 函数
// 进行基准测试
func BenchmarkItoa(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}

// go test -v -run="none" -bench="BenchmarkSprintf"
// go test -v -run="none" -bench="BenchmarkSprintf -benchtime="3s" -benchmem
// go test -v -run="none" -bench=. -benchmem

// G:\GitHub\learnGo\Go语言实战\ch09_测试和性能\listing28>go test -v -run="none" -bench=. -benchmem
// BenchmarkSprintf-4      10000000               148 ns/op              16 B/op          2 allocs/op
// BenchmarkFormat-4       30000000                43.1 ns/op             2 B/op          1 allocs/op
// BenchmarkItoa-4         30000000                45.4 ns/op             2 B/op          1 allocs/op
// PASS
// ok      _/G_/GitHub/learnGo/Go语言实战/ch09_测试和性能/listing28        6.056s
