/*
Q5. (0) 平均值
编写计算一个类型是 float64 的 slice 的平均值的代码
*/
package main

import (
	"fmt"
)

func average(xs []float64) (avg float64) {
	sum := 0.0
	switch len(xs) {
	case 0:
		avg = 0
	default:
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))
	}
	return
}

func main() {
	slice := []float64{1.0, 2.123, 3.21, 6.332, 100.123}
	fmt.Printf("avg: %v", average(slice))
}
